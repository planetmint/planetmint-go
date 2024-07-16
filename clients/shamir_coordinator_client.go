package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/planetmint/planetmint-go/config"
	"github.com/rddl-network/go-utils/tls"
)

// TODO: revert to actual rddl-claim-service client after CosmosSDK upgrade to v0.50.x
// see https://github.com/planetmint/planetmint-go/issues/384

var ShamirCoordinatorServiceClient IShamirCoordinatorClient

func lazyLoadShamirCoordinatorClient() IShamirCoordinatorClient {
	if ShamirCoordinatorServiceClient != nil {
		return ShamirCoordinatorServiceClient
	}
	cfg := config.GetConfig()
	httpsClient, err := tls.Get2WayTLSClient(cfg.CertsPath)
	if err != nil {
		defer log.Fatal("fatal error setting up mutual tls client for shamir coordinator")
	}
	ShamirCoordinatorServiceClient = NewShamirCoordinatorClient(cfg.IssuerHost, httpsClient)
	return ShamirCoordinatorServiceClient
}

func SendTokens(ctx context.Context, recipient string, amount string, asset string) (txID string, err error) {
	client := lazyLoadShamirCoordinatorClient()
	res, err := client.SendTokens(ctx, recipient, amount, asset)
	if err != nil {
		return
	}
	return res.TxID, nil
}

func ReIssueAsset(ctx context.Context, asset string, amount string) (txID string, err error) {
	client := lazyLoadShamirCoordinatorClient()
	res, err := client.ReIssueAsset(ctx, asset, amount)
	if err != nil {
		return
	}
	return res.TxID, nil
}

func IssueNFTAsset(ctx context.Context, name string, machineAddress string, domain string) (assetID string, contract string, hexTx string, err error) {
	client := lazyLoadShamirCoordinatorClient()
	res, err := client.IssueMachineNFT(ctx, name, machineAddress, domain)
	if err != nil {
		return
	}
	return res.Asset, res.Contract, res.HexTX, nil
}

type IShamirCoordinatorClient interface {
	GetMnemonics(ctx context.Context) (res MnemonicsResponse, err error)
	PostMnemonics(ctx context.Context, secret string) (err error)
	SendTokens(ctx context.Context, recipient string, amount string, asset string) (res SendTokensResponse, err error)
	ReIssueAsset(ctx context.Context, asset string, amount string) (res ReIssueResponse, err error)
	IssueMachineNFT(ctx context.Context, name string, machineAddress string, domain string) (res IssueMachineNFTResponse, err error)
}

type SendTokensRequest struct {
	Recipient string `binding:"required" json:"recipient"`
	Amount    string `binding:"required" json:"amount"`
	Asset     string `                   json:"asset"`
	ID        int    `                   json:"id"`
}

type SendTokensResponse struct {
	TxID string `binding:"required" json:"tx-id"`
}

type ReIssueRequest struct {
	Asset  string `binding:"required" json:"asset"`
	Amount string `binding:"required" json:"amount"`
	ID     int    `                   json:"id"`
}

type ReIssueResponse struct {
	TxID string `binding:"required" json:"tx-id"`
}

type MnemonicsResponse struct {
	Mnemonics []string `binding:"required" json:"mnemonics"`
	Seed      string   `binding:"required" json:"seed"`
}

type IssueMachineNFTRequest struct {
	Name           string `binding:"required" json:"name"`
	MachineAddress string `binding:"required" json:"machine-address"`
	Domain         string `binding:"required" json:"domain"`
	ID             int    `                   json:"id"`
}

type IssueMachineNFTResponse struct {
	Asset    string `binding:"required" json:"asset"`
	Contract string `binding:"required" json:"contract"`
	HexTX    string `binding:"required" json:"hex-tx"`
}

type ShamirCoordinatorClient struct {
	baseURL string
	client  *http.Client
}

func NewShamirCoordinatorClient(baseURL string, client *http.Client) *ShamirCoordinatorClient {
	if client == nil {
		client = &http.Client{}
	}
	return &ShamirCoordinatorClient{
		baseURL: baseURL,
		client:  client,
	}
}

func (scc *ShamirCoordinatorClient) GetMnemonics(ctx context.Context) (res MnemonicsResponse, err error) {
	err = scc.doRequest(ctx, http.MethodGet, scc.baseURL+"/mnemonics", nil, &res)
	return
}

func (scc *ShamirCoordinatorClient) PostMnemonics(ctx context.Context, secret string) (err error) {
	err = scc.doRequest(ctx, http.MethodPost, scc.baseURL+"/mnemonics/"+url.PathEscape(secret), nil, nil)
	return
}

func (scc *ShamirCoordinatorClient) SendTokens(ctx context.Context, recipient string, amount string, asset string) (res SendTokensResponse, err error) {
	requestBody := SendTokensRequest{
		Recipient: recipient,
		Amount:    amount,
		Asset:     asset,
	}
	err = scc.doRequest(ctx, http.MethodPost, scc.baseURL+"/send", &requestBody, &res)
	return
}

func (scc *ShamirCoordinatorClient) ReIssueAsset(ctx context.Context, asset string, amount string) (res ReIssueResponse, err error) {
	requestBody := ReIssueRequest{
		Asset:  asset,
		Amount: amount,
	}
	err = scc.doRequest(ctx, http.MethodPost, scc.baseURL+"/reissue", &requestBody, &res)
	return
}

func (scc *ShamirCoordinatorClient) IssueMachineNFT(ctx context.Context, name string, machineAddress string, domain string) (res IssueMachineNFTResponse, err error) {
	requestBody := IssueMachineNFTRequest{
		Name:           name,
		MachineAddress: machineAddress,
		Domain:         domain,
	}
	err = scc.doRequest(ctx, http.MethodPost, scc.baseURL+"/issue-machine-nft", &requestBody, &res)
	return
}

func (scc *ShamirCoordinatorClient) doRequest(ctx context.Context, method, url string, body interface{}, response interface{}) (err error) {
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

	resp, err := scc.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return &sccHTTPError{StatusCode: resp.StatusCode, Msg: strings.Join(resp.Header["Error"], "\n")}
	}

	if response != nil {
		return json.NewDecoder(resp.Body).Decode(response)
	}

	return
}

type sccHTTPError struct {
	StatusCode int
	Msg        string
}

func (e *sccHTTPError) Error() string {
	return http.StatusText(e.StatusCode) + ": " + e.Msg
}
