package main

import (
	"github.com/gin-gonic/gin"
)

type invokation struct {
	ChannelName   string `json:"channel_name"`
	ChaincodeName string `json:"chaincode_name"`
	FunctionName  string `json:"function_name"`
}

/*
API endpoints:
http://{{API_URI}}/query
{{CHANNEL_NAME}}/{{CHAINCODE_NAME}}/{{FUNCTION/NAME}}
*/
func main() {
	router := gin.Default()
	router.GET("/")
	router.Run("localhost" + "PUERTO")
}
