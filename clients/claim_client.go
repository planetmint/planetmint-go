package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/planetmint/planetmint-go/config"
)

// TODO: revert to actual rddl-claim-service client after CosmosSDK upgrade to v0.50.x
// see https://github.com/planetmint/planetmint-go/issues/384

var ClaimServiceClient IRCClient

func lazyLoad() IRCClient {
	if ClaimServiceClient != nil {
		return ClaimServiceClient
	}
	cfg := config.GetConfig()
	ClaimServiceClient = NewRCClient(cfg.ClaimHost, &http.Client{})
	return ClaimServiceClient
}

func PostClaim(ctx context.Context, beneficiary string, amount uint64, id uint64) (txID string, err error) {
	client := lazyLoad()
	res, err := client.PostClaim(ctx, PostClaimRequest{Beneficiary: beneficiary, Amount: amount, ClaimID: int(id)})
	if err != nil {
		return
	}
	return res.TxID, nil
}

type PostClaimRequest struct {
	Beneficiary string `binding:"required" json:"beneficiary"`
	Amount      uint64 `binding:"required" json:"amount"`
	ClaimID     int    `binding:"required" json:"claim-id"`
}

type PostClaimResponse struct {
	ID   int    `binding:"required" json:"id"`
	TxID string `binding:"required" json:"tx-id"`
}

type GetClaimResponse struct {
	ID           int    `binding:"required" json:"id"`
	Beneficiary  string `binding:"required" json:"beneficiary"`
	Amount       uint64 `binding:"required" json:"amount"`
	LiquidTXHash string `binding:"required" json:"liquid-tx-hash"`
	ClaimID      int    `binding:"required" json:"claim-id"`
}

type IRCClient interface {
	GetClaim(ctx context.Context, id int) (res GetClaimResponse, err error)
	PostClaim(ctx context.Context, req PostClaimRequest) (res PostClaimResponse, err error)
}

type RCClient struct {
	baseURL string
	client  *http.Client
}

func NewRCClient(baseURL string, client *http.Client) *RCClient {
	if client == nil {
		client = &http.Client{}
	}
	return &RCClient{
		baseURL: baseURL,
		client:  client,
	}
}

func (rcc *RCClient) GetClaim(ctx context.Context, id int) (res GetClaimResponse, err error) {
	err = rcc.doRequest(ctx, http.MethodGet, rcc.baseURL+"/claim/"+strconv.Itoa(id), nil, &res)
	return
}

func (rcc *RCClient) PostClaim(ctx context.Context, req PostClaimRequest) (res PostClaimResponse, err error) {
	err = rcc.doRequest(ctx, http.MethodPost, rcc.baseURL+"/claim", req, &res)
	return
}

func (rcc *RCClient) doRequest(ctx context.Context, method, url string, body interface{}, response interface{}) (err error) {
	var bodyReader io.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewBuffer(bodyBytes)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := rcc.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return &httpError{StatusCode: resp.StatusCode}
	}

	if response != nil {
		return json.NewDecoder(resp.Body).Decode(response)
	}

	return
}

type httpError struct {
	StatusCode int
}

func (e *httpError) Error() string {
	return http.StatusText(e.StatusCode)
}
