// Code generated by MockGen. DO NOT EDIT.
// Source: x/machine/keeper/register_nft.go

// Package testutil is a generated GoMock package.
package mocks

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type Body struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
}

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

// GetDoFunc fetches the mock client's `Do` func
func GetDoFunc(req *http.Request) (*http.Response, error) {
	var body Body
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		return nil, err
	}

	result := `{ "assetid": "0000000000000000000000000000000000000000000000000000000000000000"}`
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(result)),
	}
	return resp, nil
}

// Do is the mock client's `Do` func
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}
