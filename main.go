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

// TODO CHECK ENVS

func main() {
	router := gin.Default()

	router.POST("/invoke", invoke)
	router.POST("/query", query)

	router.Run("0.0.0.0:3000") // TODO MAKE PORT ENV
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
