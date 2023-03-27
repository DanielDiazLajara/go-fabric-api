package blockchain

import (
	"context"
	"fmt"
	"go-fabric-api/src/utils"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

type EventListener struct {
	ChannelName   string `json:"channel_name"`
	ChaincodeName string `json:"chaincode_name"`
}

func (listener *EventListener) Start(connection *Connection) error {
	clientConnection := connection.NewGrpcConnection()
	defer clientConnection.Close()

	id := connection.NewIdentity()
	sign := connection.NewSign()

	gateway, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	defer gateway.Close()

	network := gateway.GetNetwork(listener.ChannelName)

	// Context used for event listening
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println("\n--> Start chaincode event listening")

	events, err := network.ChaincodeEvents(ctx, listener.ChaincodeName)
	if err != nil {
		return fmt.Errorf("failed to start chaincode event listening: %w", err)
	}

	go func() {
		for event := range events {
			asset := utils.FormatJSON(event.Payload)
			fmt.Printf("\n<-- Chaincode event received: %s - %s\n", event.EventName, asset)
		}
	}()

	return nil
}
