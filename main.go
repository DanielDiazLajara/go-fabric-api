package main

import (
	"go-fabric-api/src/blockchain"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const paramsInvokeError = "expected body params: channel_name <string>, chaincode_name <string>, function_name <string>, and arguments <[]string>"

var connection blockchain.Connection

func init() {
	connection.MspID = os.Getenv("MSP_ID")
	connection.CertPath = os.Getenv("CERT_PATH")
	connection.KeyPath = os.Getenv("KEY_PATH")
	connection.TlsCertPath = os.Getenv("TLS_CERT_PATH")
	connection.PeerEndpoint = os.Getenv("PEER_ENDPOINT")
	connection.GatewayPeer = os.Getenv("PEER_GATEWAY")
}

/*
	mspID        = "Org1MSP"
	cryptoPath   = "../../test-network/organizations/peerOrganizations/org1.example.com"
	certPath     = cryptoPath + "/users/User1@org1.example.com/msp/signcerts/cert.pem"
	keyPath      = cryptoPath + "/users/User1@org1.example.com/msp/keystore/"
	tlsCertPath  = cryptoPath + "/peers/peer0.org1.example.com/tls/ca.crt"
	peerEndpoint = "localhost:7051"
	gatewayPeer  = "peer0.org1.example.com"
*/

func main() {
	router := gin.Default()

	router.POST("/invoke", invoke)
	router.POST("/query", query)

	router.Run("localhost:3000") // TODO MAKE PORT ENV
}

func invoke(c *gin.Context) {
	var invoke blockchain.Invokation
	if err := c.BindJSON(&invoke); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, paramsInvokeError)
		return
	}
	result, err := invoke.Execute(&connection)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "transaction invoke failed: "+err.Error())
		return
	}
	c.IndentedJSON(http.StatusCreated, result)
}

func query(c *gin.Context) {
	var query blockchain.Query
	if err := c.BindJSON(&query); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, paramsInvokeError)
		return
	}
	result, err := query.Execute(&connection)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "transaction invoke failed: "+err.Error())
		return
	}
	c.IndentedJSON(http.StatusCreated, result)
}
