/*
Copyright 2022 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package blockchain

import (
	"crypto/x509"
	"fmt"
	"os"
	"path"

	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Connection struct {
	mspID        string
	cryptoPath   string
	certPath     string
	keyPath      string
	tlsCertPath  string
	peerEndpoint string
	gatewayPeer  string
}

// newGrpcConnection creates a gRPC connection to the Gateway server.
func (con *Connection) NewGrpcConnection() *grpc.ClientConn {
	certificate, err := loadCertificate(con.tlsCertPath)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, con.gatewayPeer)

	connection, err := grpc.Dial(con.peerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}

// newIdentity creates a client identity for this Gateway connection using an X.509 certificate.
func (con *Connection) NewIdentity() *identity.X509Identity {
	certificate, err := loadCertificate(con.certPath)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(con.mspID, certificate)
	if err != nil {
		panic(err)
	}

	return id
}

func loadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}

// newSign creates a function that generates a digital signature from a message digest using a private key.
func (con *Connection) NewSign() identity.Sign {
	files, err := os.ReadDir(con.keyPath)
	if err != nil {
		panic(fmt.Errorf("failed to read private key directory: %w", err))
	}
	privateKeyPEM, err := os.ReadFile(path.Join(con.keyPath, files[0].Name()))

	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}
