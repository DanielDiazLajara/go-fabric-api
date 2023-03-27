package blockchain

import (
	"fmt"
	"go-fabric-api/src/utils"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

type Query struct {
	ChannelName   string   `json:"channel_name"`
	ChaincodeName string   `json:"chaincode_name"`
	FunctionName  string   `json:"function_name"`
	Arguments     []string `json:"arguments"`
}

// Evaluate a transaction by assetID to query ledger state.
func (query *Query) Execute(connection *Connection) (string, error) {
	// The gRPC client connection should be shared by all Gateway connections to this endpoint
	clientConnection := connection.NewGrpcConnection()
	defer clientConnection.Close()

	id := connection.NewIdentity()
	sign := connection.NewSign()

	// Create a Gateway connection for a specific client identity
	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		return "", fmt.Errorf("failed to connect to network: %w", err)
	}
	defer gw.Close()

	chaincodeName := query.ChaincodeName

	channelName := query.ChannelName

	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)
	fmt.Printf("\n--> Evaluating Transaction in channel " + query.ChannelName + ", chaincode " + query.ChaincodeName + ", function " + query.FunctionName + "\n")

	evaluateResult, err := contract.EvaluateTransaction(query.FunctionName, query.Arguments...)
	if err != nil {
		return "", fmt.Errorf("failed to evaluate transaction: %w", err)
	}
	result := utils.FormatJSON(evaluateResult)

	return string(result), nil

}
