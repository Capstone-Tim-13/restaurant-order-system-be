package database

import (
	"capstone/config"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

func MidtransSnapClient(config config.DatabaseConfig) snap.Client {
	var snapClient snap.Client
	snapClient.New(config.SERVER_KEY_MT, midtrans.Sandbox)

	return snapClient
}

func MidtransCoreAPIClient(config config.DatabaseConfig) coreapi.Client {
	var coreAPIClient coreapi.Client
	coreAPIClient.New(config.SERVER_KEY_MT, midtrans.Sandbox)

	return coreAPIClient
}