package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

var (
	mspID        = ""
	certPath     = ""
	keyPath      = ""
	tlsCertPath  = ""
	peerEndpoint = ""
	gatewayPeer  = ""
)

func init() {
	mspID = os.Getenv("MSP_ID")
	certPath = os.Getenv("CERT_PATH")
	keyPath = os.Getenv("KEY_PATH")
	tlsCertPath = os.Getenv("TLS_CERT_PATH")
	peerEndpoint = os.Getenv("PEER_ENDPOINT")
	gatewayPeer = os.Getenv("PEER_GATEWAY")
	/*
		mspID        = "Org1MSP"
		cryptoPath   = "../../test-network/organizations/peerOrganizations/org1.example.com"
		certPath     = cryptoPath + "/users/User1@org1.example.com/msp/signcerts/cert.pem"
		keyPath      = cryptoPath + "/users/User1@org1.example.com/msp/keystore/"
		tlsCertPath  = cryptoPath + "/peers/peer0.org1.example.com/tls/ca.crt"
		peerEndpoint = "localhost:7051"
		gatewayPeer  = "peer0.org1.example.com"
	*/
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
