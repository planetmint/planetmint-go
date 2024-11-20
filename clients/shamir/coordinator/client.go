package coordinator

import (
	"context"
	"log"

	"github.com/planetmint/planetmint-go/config"
	"github.com/rddl-network/go-utils/tls"
	"github.com/rddl-network/shamir-coordinator-service/client"
)

var ShamirCoordinatorServiceClient client.IShamirCoordinatorClient

func lazyLoad() client.IShamirCoordinatorClient {
	if ShamirCoordinatorServiceClient != nil {
		return ShamirCoordinatorServiceClient
	}
	cfg := config.GetConfig()
	httpsClient, err := tls.Get2WayTLSClient(cfg.CertsPath)
	if err != nil {
		defer log.Fatal("fatal error setting up mutual tls client for shamir coordinator")
	}
	ShamirCoordinatorServiceClient = client.NewShamirCoordinatorClient(cfg.IssuerHost, httpsClient)
	return ShamirCoordinatorServiceClient
}

func SendTokens(ctx context.Context, recipient string, amount string, asset string) (txID string, err error) {
	client := lazyLoad()
	res, err := client.SendTokens(ctx, recipient, amount, asset)
	if err != nil {
		return
	}
	return res.TxID, nil
}

func ReIssueAsset(ctx context.Context, asset string, amount string) (txID string, err error) {
	client := lazyLoad()
	res, err := client.ReIssueAsset(ctx, asset, amount)
	if err != nil {
		return
	}
	return res.TxID, nil
}

func IssueNFTAsset(ctx context.Context, name string, machineAddress string, domain string) (assetID string, contract string, hexTx string, err error) {
	client := lazyLoad()
	res, err := client.IssueMachineNFT(ctx, name, machineAddress, domain)
	if err != nil {
		return
	}
	return res.Asset, res.Contract, res.HexTX, nil
}
