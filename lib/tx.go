package lib

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"syscall"

	"github.com/cometbft/cometbft/crypto"
	comethttp "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/planetmint/planetmint-go/lib/trustwallet"
)

var (
	ErrTypeAssertionFailed = errors.New("type assertion failed")
	LibSyncAccess          sync.Mutex
)

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

func getClientContextAndTxFactory(fromAddress sdk.AccAddress, withoutFee bool) (clientCtx client.Context, txf tx.Factory, err error) {
	clientCtx = GetConfig().clientCtx
	// at least we need an account retriever
	// it would be better to check for an empty client context, but that does not work at the moment
	if clientCtx.AccountRetriever == nil {
		clientCtx, err = getClientContext(fromAddress)
		if err != nil {
			return
		}
	}
	record, err := clientCtx.Keyring.KeyByAddress(fromAddress)
	if err != nil {
		return
	}
	// name and address of private key with which to sign
	clientCtx = clientCtx.
		WithFromAddress(fromAddress).
		WithFromName(record.Name)

	accountNumber, sequence, err := getAccountNumberAndSequence(clientCtx)
	if err != nil {
		return
	}
	gasPrice := "0.000005"
	if withoutFee {
		gasPrice = "0.0"
	}
	txf = getTxFactoryWithAccountNumberAndSequence(clientCtx, accountNumber, sequence, gasPrice)
	return
}

func getTxFactoryWithAccountNumberAndSequence(clientCtx client.Context, accountNumber, sequence uint64, gasPrice string) (txf tx.Factory) {
	return tx.Factory{}.
		WithAccountNumber(accountNumber).
		WithAccountRetriever(clientCtx.AccountRetriever).
		WithChainID(clientCtx.ChainID).
		WithFeeGranter(clientCtx.FeeGranter).
		WithGas(GetConfig().txGas).
		WithGasPrices(gasPrice + GetConfig().feeDenom).
		WithKeybase(clientCtx.Keyring).
		WithSequence(sequence).
		WithTxConfig(clientCtx.TxConfig)
}

func getClientContext(fromAddress sdk.AccAddress) (clientCtx client.Context, err error) {
	encodingConfig := GetConfig().encodingConfig

	rootDir := GetConfig().rootDir
	input := os.Stdin
	codec := encodingConfig.Marshaler
	keyringOptions := []keyring.Option{}

	keyring, err := keyring.New("lib", keyring.BackendTest, rootDir, input, codec, keyringOptions...)
	if err != nil {
		return
	}

	record, err := keyring.KeyByAddress(fromAddress)
	if err != nil {
		return
	}

	remote := GetConfig().rpcEndpoint
	wsClient, err := comethttp.New(remote, "/websocket")
	if err != nil {
		return
	}

	var output bytes.Buffer

	clientCtx = client.Context{
		AccountRetriever:  authtypes.AccountRetriever{},
		BroadcastMode:     "sync",
		ChainID:           GetConfig().chainID,
		Client:            wsClient,
		Codec:             codec,
		From:              fromAddress.String(),
		FromAddress:       fromAddress,
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

func isMachineAttestationMsg(msgs ...sdk.Msg) (isMachineAttestation bool) {
	if len(msgs) != 1 {
		return
	}
	if sdk.MsgTypeURL(msgs[0]) == "/planetmintgo.machine.MsgAttestMachine" {
		isMachineAttestation = true
	}
	return
}

// BuildUnsignedTx builds a transaction to be signed given a set of messages.
// Once created, the fee, memo, and messages are set.
func BuildUnsignedTx(fromAddress sdk.AccAddress, msgs ...sdk.Msg) (txJSON string, err error) {
	withoutFee := isMachineAttestationMsg(msgs...)
	clientCtx, txf, err := getClientContextAndTxFactory(fromAddress, withoutFee)
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

func broadcastTx(clientCtx client.Context, txf tx.Factory, msgs ...sdk.Msg) (out *bytes.Buffer, err error) {
	err = tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msgs...)
	if err != nil {
		return
	}
	return writeClientCtxOutputToBuffer(clientCtx)
}

// BroadcastTxWithFileLock broadcasts a transaction via gRPC and synchronises requests via a file lock.
func BroadcastTxWithFileLock(fromAddress sdk.AccAddress, msgs ...sdk.Msg) (out *bytes.Buffer, err error) {
	LibSyncAccess.Lock()
	defer LibSyncAccess.Unlock()
	// open and lock file, if it exists
	file, err := openSequenceFile(fromAddress)
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
	withoutFee := isMachineAttestationMsg(msgs...)
	clientCtx, txf, err := getClientContextAndTxFactory(fromAddress, withoutFee)
	if err != nil {
		return
	}

	sequenceFromFile, errFile := getSequenceFromFile(file)
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
	if GetConfig().serialPort != "" {
		out, err = broadcastTxWithTrustWalletSignature(clientCtx, txf, msgs...)
	} else {
		out, err = broadcastTx(clientCtx, txf, msgs...)
	}
	if err != nil {
		return
	}

	txResponse, err := GetTxResponseFromOut(out)
	if err != nil {
		return
	}

	// Only increase counter if broadcast was successful
	if txResponse.Code != 0 {
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

func broadcastTxWithTrustWalletSignature(clientCtx client.Context, txf tx.Factory, msgs ...sdk.Msg) (out *bytes.Buffer, err error) {
	txBuilder, err := txf.BuildUnsignedTx(msgs...)
	if err != nil {
		return
	}

	if err = signWithTrustWallet(txf, clientCtx, txBuilder); err != nil {
		return
	}

	txBytes, err := clientCtx.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return
	}

	res, err := clientCtx.BroadcastTx(txBytes)
	if err != nil {
		return
	}

	if err = clientCtx.PrintProto(res); err != nil {
		return
	}

	return writeClientCtxOutputToBuffer(clientCtx)
}

func writeClientCtxOutputToBuffer(clientCtx client.Context) (out *bytes.Buffer, err error) {
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

	// Make a copy because we `defer output.Reset()`
	out = &bytes.Buffer{}
	// This is still copying references: *out = *output
	// Make a real copy: https://stackoverflow.com/a/69758157
	out.Write(output.Bytes())
	return
}

func signWithTrustWallet(txf tx.Factory, clientCtx client.Context, txBuilder client.TxBuilder) error {
	connector, err := trustwallet.NewTrustWalletConnector(GetConfig().serialPort)
	if err != nil {
		return err
	}

	kb, err := clientCtx.Keyring.Key(clientCtx.FromName)
	if err != nil {
		return err
	}

	pubkey, err := kb.GetPubKey()
	if err != nil {
		return err
	}

	signMode := txf.SignMode()
	if signMode == signing.SignMode_SIGN_MODE_UNSPECIFIED {
		// use the SignModeHandler's default mode if unspecified
		signMode = clientCtx.TxConfig.SignModeHandler().DefaultMode()
	}

	signerData := authsigning.SignerData{
		ChainID:       txf.ChainID(),
		AccountNumber: txf.AccountNumber(),
		Sequence:      txf.Sequence(),
		PubKey:        pubkey,
		Address:       sdk.AccAddress(pubkey.Address()).String(),
	}

	sigData := signing.SingleSignatureData{
		SignMode:  signMode,
		Signature: nil,
	}

	sig := signing.SignatureV2{
		PubKey:   pubkey,
		Data:     &sigData,
		Sequence: txf.Sequence(),
	}

	if err := txBuilder.SetSignatures(sig); err != nil {
		return err
	}

	bytesToSign, err := clientCtx.TxConfig.SignModeHandler().GetSignBytes(signMode, signerData, txBuilder.GetTx())
	if err != nil {
		return err
	}

	hashBytesToSign := crypto.Sha256(bytesToSign)
	hexHash := hex.EncodeToString(hashBytesToSign)

	hexSig, err := connector.SignHashWithPlanetmint(hexHash)
	if err != nil {
		return err
	}

	signature, err := hex.DecodeString(hexSig)
	if err != nil {
		return err
	}

	sigData = signing.SingleSignatureData{
		SignMode:  signMode,
		Signature: signature,
	}
	sig = signing.SignatureV2{
		PubKey:   pubkey,
		Data:     &sigData,
		Sequence: txf.Sequence(),
	}

	if err = txBuilder.SetSignatures(sig); err != nil {
		return fmt.Errorf("unable to set signatures on payload: %w", err)
	}

	// Run optional preprocessing if specified. By default, this is unset
	// and will return nil.
	return txf.PreprocessTx(clientCtx.FromName, txBuilder)
}
