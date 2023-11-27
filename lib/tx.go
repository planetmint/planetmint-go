package lib

import (
	"bufio"
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	comethttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Result defines a generic way to receive responses from the RPC endpoint.
type Result struct {
	Info map[string]interface{} `json:"info" mapstructure:"info"`
}

func init() {
	GetConfig()
}

func getAccountNumberAndSequence(goCtx context.Context, address sdk.AccAddress) (accountNumber, sequence uint64, err error) {
	url := fmt.Sprintf("%s/cosmos/auth/v1beta1/account_info/%s", libConfig.APIEndpoint, address.String())
	req, err := http.NewRequestWithContext(goCtx, http.MethodGet, url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var result Result
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return
	}

	accountNumber, err = strconv.ParseUint(result.Info["account_number"].(string), 10, 64)
	if err != nil {
		return
	}
	sequence, err = strconv.ParseUint(result.Info["sequence"].(string), 10, 64)
	if err != nil {
		return
	}
	return
}

func getClientContextAndTxFactory(goCtx context.Context, address sdk.AccAddress) (clientCtx client.Context, txf tx.Factory, err error) {
	clientCtx, err = getClientContext(address)
	if err != nil {
		return
	}
	accountNumber, sequence, err := getAccountNumberAndSequence(goCtx, clientCtx.FromAddress)
	if err != nil {
		return
	}
	txf = getTxFactoryWithAccountNumberAndSequence(clientCtx, accountNumber, sequence)
	return
}

func getClientContextAndTxFactoryWithAccountNumberAndSequence(address sdk.AccAddress, accountNumber, sequence uint64) (clientCtx client.Context, txf tx.Factory, err error) {
	clientCtx, err = getClientContext(address)
	if err != nil {
		return
	}
	txf = getTxFactoryWithAccountNumberAndSequence(clientCtx, accountNumber, sequence)
	return
}

func getTxFactoryWithAccountNumberAndSequence(clientCtx client.Context, accountNumber, sequence uint64) (txf tx.Factory) {
	return tx.Factory{}.
		WithAccountNumber(accountNumber).
		WithChainID(clientCtx.ChainID).
		WithGas(200000).
		WithGasPrices("0.000005plmnt").
		WithKeybase(clientCtx.Keyring).
		WithSequence(sequence).
		WithTxConfig(clientCtx.TxConfig)
}

func getClientContext(address sdk.AccAddress) (clientCtx client.Context, err error) {
	encodingConfig := GetConfig().EncodingConfig

	rootDir := GetConfig().RootDir
	input := os.Stdin
	codec := encodingConfig.Marshaler
	keyringOptions := []keyring.Option{}

	keyring, err := keyring.New("lib", keyring.BackendTest, rootDir, input, codec, keyringOptions...)
	if err != nil {
		return
	}

	record, err := keyring.KeyByAddress(address)
	if err != nil {
		return
	}

	remote := GetConfig().RPCEndpoint
	wsClient, err := comethttp.New(remote, "/websocket")
	if err != nil {
		return
	}

	var output bytes.Buffer

	clientCtx = client.Context{
		BroadcastMode:  "sync",
		ChainID:        "planetmint-testnet-1",
		Client:         wsClient,
		Codec:          codec,
		From:           address.String(),
		FromAddress:    address,
		FromName:       record.Name,
		HomeDir:        rootDir,
		Input:          input,
		Keyring:        keyring,
		KeyringDir:     rootDir,
		KeyringOptions: keyringOptions,
		NodeURI:        remote,
		Offline:        true,
		Output:         &output,
		SkipConfirm:    true,
		TxConfig:       encodingConfig.TxConfig,
	}

	return
}

// BuildUnsignedTx builds a transaction to be signed given a set of messages.
// Once created, the fee, memo, and messages are set.
func BuildUnsignedTx(goCtx context.Context, address sdk.AccAddress, msgs ...sdk.Msg) (txJSON string, err error) {
	clientCtx, txf, err := getClientContextAndTxFactory(goCtx, address)
	if err != nil {
		return
	}
	txBuilder, err := txf.BuildUnsignedTx(msgs...)
	if err != nil {
		return
	}
	// Generate a JSON string.
	txJSONBytes, err := clientCtx.TxConfig.TxJSONEncoder()(txBuilder.GetTx())
	if err != nil {
		return
	}
	txJSON = string(txJSONBytes)
	return
}

// BroadcastTx broadcasts a transaction via RPC.
func BroadcastTx(goCtx context.Context, address sdk.AccAddress, msgs ...sdk.Msg) (broadcastTxResponseJSON string, err error) {
	clientCtx, txf, err := getClientContextAndTxFactory(goCtx, address)
	if err != nil {
		return
	}
	broadcastTxResponseJSON, err = broadcastTx(clientCtx, txf, msgs...)
	return
}

func broadcastTxWithAccountNumberAndSequence(address sdk.AccAddress, accountNumber, sequence uint64, msgs ...sdk.Msg) (broadcastTxResponseJSON string, err error) {
	clientCtx, txf, err := getClientContextAndTxFactoryWithAccountNumberAndSequence(address, accountNumber, sequence)
	if err != nil {
		return
	}
	broadcastTxResponseJSON, err = broadcastTx(clientCtx, txf, msgs...)
	return
}

func broadcastTx(clientCtx client.Context, txf tx.Factory, msgs ...sdk.Msg) (broadcastTxResponseJSON string, err error) {
	err = tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msgs...)
	output, ok := clientCtx.Output.(*bytes.Buffer)
	if !ok {
		err = errors.New("type assertion failed")
		return
	}
	broadcastTxResponseJSON = output.String()
	return
}

// BroadcastTxWithFileLock broadcasts a transaction via gRPC and synchronises requests via a file lock.
func BroadcastTxWithFileLock(goCtx context.Context, address sdk.AccAddress, msgs ...sdk.Msg) (broadcastTxResponseJSON string, err error) {
	usr, err := user.Current()
	if err != nil {
		return
	}
	homeDir := usr.HomeDir

	addrHex := hex.EncodeToString(address)
	filename := filepath.Join(GetConfig().RootDir, addrHex+".sequence")

	// Expand tilde to user's home directory.
	if filename == "~" {
		filename = homeDir
	} else if strings.HasPrefix(filename, "~/") {
		filename = filepath.Join(homeDir, filename[2:])
	}

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	// Get file lock.
	err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX)
	if err != nil {
		return
	}
	defer func() {
		if err := syscall.Flock(int(file.Fd()), syscall.LOCK_UN); err != nil {
			return
		}
	}()

	var sequenceString string
	lineCount := int64(0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sequenceString = scanner.Text()
		lineCount++
	}
	err = scanner.Err()
	if err != nil {
		return
	}

	// Get sequence number from chain.
	accountNumber, sequence, err := getAccountNumberAndSequence(goCtx, address)
	if err != nil {
		return
	}

	if lineCount == 0 {
		// File does not exist yet.
		sequenceString = strconv.FormatUint(sequence, 10)
	} else if lineCount != 1 {
		err = errors.New("Malformed " + filename + ": wrong number of lines")
		return
	}

	sequenceCount, err := strconv.ParseUint(sequenceString, 10, 64)
	if err != nil {
		return
	}

	// Sequence number on chain is bigger than in text file.
	// Someone manually sent a transaction from our account?
	if sequence > sequenceCount {
		sequenceCount = sequence
	}

	broadcastTxResponseJSON, err = broadcastTxWithAccountNumberAndSequence(address, accountNumber, sequenceCount, msgs...)
	if err != nil {
		return
	}

	// Increase counter for next round.
	sequenceCount++

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return
	}

	_, err = file.WriteString(strconv.FormatUint(sequenceCount, 10) + "\n")

	return
}
