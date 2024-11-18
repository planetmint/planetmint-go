package clients

import (
	"context"
	"net/http"

	"github.com/planetmint/planetmint-go/config"
	"github.com/rddl-network/rddl-claim-service/client"
	"github.com/rddl-network/rddl-claim-service/types"
)

var ClaimServiceClient client.IRCClient

func lazyLoad() client.IRCClient {
	if ClaimServiceClient != nil {
		return ClaimServiceClient
	}
	cfg := config.GetConfig()
	ClaimServiceClient = client.NewRCClient(cfg.ClaimHost, &http.Client{})
	return ClaimServiceClient
}

func PostClaim(ctx context.Context, beneficiary string, amount uint64, id uint64) (txID string, err error) {
	client := lazyLoad()
	res, err := client.PostClaim(ctx, types.PostClaimRequest{Beneficiary: beneficiary, Amount: amount, ClaimID: int(id)})
	if err != nil {
		return
	}
	return res.TxID, nil
}
