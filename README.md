# go-fabric-api
Plug-and-play Go-based API for Hyperledger Fabric

## Configuration
Create and fill a 📄.env file containing the following variables:
- MSP_ID=Org1MSP
- CERT_PATH=./organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/signcerts/cert.pem
- KEY_PATH=./organizations/peerOrganizations/org1.example.com/users/User1@org1.example.com/msp/keystore/
- TLS_CERT_PATH=./organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
- PEER_ENDPOINT=localhost:7051
- PEER_GATEWAY=peer0.org1.example.com

📄 **Note:** _Should you deploy the fabric test network, you can simply copy and paste the 📁organizations folder to the go-fabric-api root folder, so you can just use the default variables. For a custom deployment, change these variables._

## Deployment
```bash
# Container
docker compose up -d (--build)

# Go
go run ./main.go

```

## Endpoints

🔵 **POST** /invoke
- Description: Executes a new transaction in the ledger, altering the world state (WS).
- Body:
```json
{
    "channel_name":"string",
    "chaincode_name":"string",
    "function_name":"string",
    "arguments":["string","string"]
}

```

🔵 **POST** /query
- Description: simulates a transaction, with no changes in either the WS or the ledger.
- Body:
```json
{
    "channel_name":"string",
    "chaincode_name":"string",
    "function_name":"string",
    "arguments":["string","string"]
}

```
