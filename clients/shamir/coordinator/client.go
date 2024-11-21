package coordinator

import (
	"context"
	"errors"

	"github.com/planetmint/planetmint-go/config"
	"github.com/rddl-network/go-utils/tls"
	"github.com/rddl-network/shamir-coordinator-service/client"
)

var (
	SCClient client.ISCClient
)

func lazyLoad() client.ISCClient {
	if SCClient != nil {
		return SCClient
	}
	cfg := config.GetConfig()
	httpsClient, err := tls.Get2WayTLSClient(cfg.CertsPath)
	if err != nil {
		err := errors.New("fatal error setting up mutual tls client for shamir coordinator")
		panic(err)
	}
	SCClient = client.NewSCClient(cfg.IssuerHost, httpsClient)
	return SCClient
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
