package lib

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
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
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

var ErrTypeAssertionFailed = errors.New("type assertion failed")

func init() {
	GetConfig()
}

func getAccountNumberAndSequence(clientCtx client.Context) (accountNumber, sequence uint64, err error) {
	account, err := clientCtx.AccountRetriever.GetAccount(clientCtx, clientCtx.FromAddress)
	if err != nil {
		return
	}
	accountNumber = account.GetAccountNumber()
	sequence = account.GetSequence()
	return
}

func getClientContextAndTxFactory(address sdk.AccAddress) (clientCtx client.Context, txf tx.Factory, err error) {
	clientCtx = GetConfig().ClientCtx
	// at least we need an account retriever
	// it would be better to check for an empty client context, but that does not work at the moment
	if clientCtx.AccountRetriever == nil {
		clientCtx, err = getClientContext(address)
		if err != nil {
			return
		}
	}
	accountNumber, sequence, err := getAccountNumberAndSequence(clientCtx)
	if err != nil {
		return
	}
	txf = getTxFactoryWithAccountNumberAndSequence(clientCtx, accountNumber, sequence)
	return
}

func getTxFactoryWithAccountNumberAndSequence(clientCtx client.Context, accountNumber, sequence uint64) (txf tx.Factory) {
	return tx.Factory{}.
		WithAccountNumber(accountNumber).
		WithAccountRetriever(clientCtx.AccountRetriever).
		WithChainID(clientCtx.ChainID).
		WithGas(200000).
		WithGasPrices("0.000005" + GetConfig().FeeDenom).
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
		AccountRetriever:  authtypes.AccountRetriever{},
		BroadcastMode:     "sync",
		ChainID:           GetConfig().ChainID,
		Client:            wsClient,
		Codec:             codec,
		From:              address.String(),
		FromAddress:       address,
		FromName:          record.Name,
		HomeDir:           rootDir,
		Input:             input,
		InterfaceRegistry: encodingConfig.InterfaceRegistry,
		Keyring:           keyring,
		KeyringDir:        rootDir,
		KeyringOptions:    keyringOptions,
		NodeURI:           remote,
		Offline:           true,
		Output:            &output,
		SkipConfirm:       true,
		TxConfig:          encodingConfig.TxConfig,
	}

	return
}

// BuildUnsignedTx builds a transaction to be signed given a set of messages.
// Once created, the fee, memo, and messages are set.
func BuildUnsignedTx(address sdk.AccAddress, msgs ...sdk.Msg) (txJSON string, err error) {
	clientCtx, txf, err := getClientContextAndTxFactory(address)
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
func BroadcastTx(address sdk.AccAddress, msgs ...sdk.Msg) (broadcastTxResponseJSON string, err error) {
	clientCtx, txf, err := getClientContextAndTxFactory(address)
	if err != nil {
		return
	}
	broadcastTxResponseJSON, err = broadcastTx(clientCtx, txf, msgs...)
	return
}

func broadcastTx(clientCtx client.Context, txf tx.Factory, msgs ...sdk.Msg) (broadcastTxResponseJSON string, err error) {
	err = tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msgs...)
	if err != nil {
		return
	}
	output, ok := clientCtx.Output.(*bytes.Buffer)
	if !ok {
		err = ErrTypeAssertionFailed
		return
	}
	defer output.Reset()

	result := make(map[string]interface{})
	err = json.Unmarshal(output.Bytes(), &result)
	if err != nil {
		return
	}
	code, ok := result["code"].(float64)
	if !ok {
		err = ErrTypeAssertionFailed
		return
	}
	if code != 0 {
		err = errors.New(output.String())
		return
	}

	broadcastTxResponseJSON = output.String()
	return
}
func getSequenceFromFile(seqFile *os.File, filename string) (sequence uint64, err error) {
	var sequenceString string
	lineCount := int64(0)
	scanner := bufio.NewScanner(seqFile)
	for scanner.Scan() {
		sequenceString = scanner.Text()
		lineCount++
	}
	err = scanner.Err()
	if err != nil {
		return
	}
	if lineCount == 0 {
		err = errors.New("Sequence file empty " + filename + ": no lines")
		return
	} else if lineCount != 1 {
		err = errors.New("Malformed " + filename + ": wrong number of lines")
		return
	}
	sequence, err = strconv.ParseUint(sequenceString, 10, 64)
	if err != nil {
		return
	}
	return
}

func getSequenceFromChain(clientCtx client.Context) (sequence uint64, err error) {
	// Get sequence number from chain.
	account, err := clientCtx.AccountRetriever.GetAccount(clientCtx, clientCtx.FromAddress)
	if err != nil {
		return
	}
	sequence = account.GetSequence()
	return
}

// BroadcastTxWithFileLock broadcasts a transaction via gRPC and synchronises requests via a file lock.
func BroadcastTxWithFileLock(address sdk.AccAddress, msgs ...sdk.Msg) (broadcastTxResponseJSON string, err error) {
	// open and lock file, if it exists
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

	// get basic chain information
	clientCtx, txf, err := getClientContextAndTxFactory(address)
	if err != nil {
		return
	}

	sequenceFromFile, errFile := getSequenceFromFile(file, filename)
	sequenceFromChain, errChain := getSequenceFromChain(clientCtx)

	var sequence uint64
	if errFile != nil && errChain != nil {
		err = errors.New("unable to determine sequence number")
		return
	}
	sequence = sequenceFromChain
	if sequenceFromFile > sequenceFromChain {
		sequence = sequenceFromFile
	}

	// Set new sequence number
	txf = txf.WithSequence(sequence)
	broadcastTxResponseJSON, err = broadcastTx(clientCtx, txf, msgs...)
	if err != nil {
		return
	}

	// Increase counter for next round.
	sequence++

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return
	}

	_, err = file.WriteString(strconv.FormatUint(sequence, 10) + "\n")

	return
}
