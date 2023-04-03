package main

import (
	"go-fabric-api/src/blockchain"
	"go-fabric-api/src/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const paramsInvokeError = "expected body params: channel_name <string>, chaincode_name <string>, function_name <string>, and arguments <[]string>"

var connection blockchain.Connection
var port int

func init() {
	var err error
	// Checking env content
	if connection.MspID, err = utils.GetenvStr("MSP_ID"); err != nil {
		panic(err.Error())
	}
	if connection.CertPath, err = utils.GetenvStr("CERT_PATH"); err != nil {
		panic(err.Error())
	}
	if connection.KeyPath, err = utils.GetenvStr("KEY_PATH"); err != nil {
		panic(err.Error())
	}
	if connection.TlsCertPath, err = utils.GetenvStr("TLS_CERT_PATH"); err != nil {
		panic(err.Error())
	}
	if connection.PeerEndpoint, err = utils.GetenvStr("PEER_ENDPOINT"); err != nil {
		panic(err.Error())
	}
	if connection.GatewayPeer, err = utils.GetenvStr("PEER_GATEWAY"); err != nil {
		panic(err.Error())
	}
	if port, err = utils.GetenvInt("PORT"); err != nil {
		panic(err.Error())
	}
	// Check files existance
	if !utils.FileExists(connection.CertPath) {
		panic(connection.CertPath + " file does not exist")
	}
	if !utils.FileExists(connection.KeyPath) {
		panic(connection.KeyPath + " file does not exist")
	}
	if !utils.FileExists(connection.TlsCertPath) {
		panic(connection.TlsCertPath + " file does not exist")
	}
}

func main() {
	router := gin.Default()

	router.POST("/invoke", invoke)
	router.POST("/query", query)

	router.Run("0.0.0.0:" + strconv.Itoa(port)) // TODO MAKE PORT ENV
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
