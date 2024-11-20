package coordinator

import (
	"context"
	"errors"

	"github.com/planetmint/planetmint-go/config"
	"github.com/rddl-network/go-utils/tls"
	"github.com/rddl-network/shamir-coordinator-service/client"
)

var (
	ShamirCoordinatorServiceClient client.IShamirCoordinatorClient
)

func lazyLoad() client.IShamirCoordinatorClient {
	if ShamirCoordinatorServiceClient != nil {
		return ShamirCoordinatorServiceClient
	}
	cfg := config.GetConfig()
	httpsClient, err := tls.Get2WayTLSClient(cfg.CertsPath)
	if err != nil {
		err := errors.New("fatal error setting up mutual tls client for shamir coordinator")
		panic(err)
	}
	ShamirCoordinatorServiceClient = client.NewShamirCoordinatorClient(cfg.IssuerHost, httpsClient)
	return ShamirCoordinatorServiceClient
}

func SendTokens(ctx context.Context, recipient string, amount string, asset string) (txID string, err error) {
	client := lazyLoad()
	resp, err := client.SendTokens(ctx, recipient, amount, asset)
	if err != nil {
		return
	}
	return resp.TxID, nil
}

func ReIssueAsset(ctx context.Context, asset string, amount string) (txID string, err error) {
	client := lazyLoad()
	resp, err := client.ReIssueAsset(ctx, asset, amount)
	if err != nil {
		return
	}
	return resp.TxID, nil
}

func IssueNFTAsset(ctx context.Context, name string, machineAddress string, domain string) (assetID string, contract string, hexTx string, err error) {
	client := lazyLoad()
	resp, err := client.IssueMachineNFT(ctx, name, machineAddress, domain)
	if err != nil {
		return
	}
	return resp.Asset, resp.Contract, resp.HexTX, nil
}
