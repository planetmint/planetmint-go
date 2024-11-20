package claim

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/planetmint/planetmint-go/config"
	"github.com/rddl-network/rddl-claim-service/client"
	"github.com/rddl-network/rddl-claim-service/types"
)

var (
	RCClient client.IRCClient
)

func lazyLoad() client.IRCClient {
	if RCClient != nil {
		return RCClient
	}
	cfg := config.GetConfig()
	RCClient = client.NewRCClient(cfg.ClaimHost, &http.Client{})
	return RCClient
}

func PostClaim(ctx context.Context, beneficiary string, amount uint64, id uint64) (txID string, err error) {
	if beneficiary == "" {
		return txID, errors.New("beneficiary cannot be empty")
	}

	if amount == 0 {
		return txID, errors.New("amount must be greater than 0")
	}

	req := types.PostClaimRequest{
		Beneficiary: beneficiary,
		Amount:      amount,
		ClaimID:     int(id),
	}

	client := lazyLoad()
	resp, err := client.PostClaim(ctx, req)
	if err != nil {
		return txID, fmt.Errorf("failed to post claim: %w", err)
	}
	return resp.TxID, nil
}
