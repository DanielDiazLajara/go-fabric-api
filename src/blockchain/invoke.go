package blockchain

import (
	"fmt"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

type Invokation struct {
	ChannelName   string   `json:"channel_name"`
	ChaincodeName string   `json:"chaincode_name"`
	FunctionName  string   `json:"function_name"`
	Arguments     []string `json:"arguments"`
}

func (inv *Invokation) Execute(connection *Connection) (string, error) {
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

	network := gateway.GetNetwork(inv.ChannelName)
	contract := network.GetContract(inv.ChaincodeName)
	fmt.Printf("\n--> Submit transaction to channel " + inv.ChannelName + ", chaincode " + inv.ChaincodeName + ", function " + inv.FunctionName + "\n")

	payload, commit, err := contract.SubmitAsync(inv.FunctionName, client.WithArguments(inv.Arguments...))
	if err != nil {
		return "", fmt.Errorf("failed to submit transaction: %w", err)
	}

	status, err := commit.Status()
	if err != nil {
		return "", fmt.Errorf("failed to get transaction commit status: %w", err)
	}

	if !status.Successful {
		return "", fmt.Errorf("failed to commit transaction with status code %v", status.Code)
	}

	return string(payload[:]), nil
}
