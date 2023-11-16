package lib

import (
	"context"
	"encoding/hex"
	"errors"
	"path/filepath"

	"github.com/99designs/keyring"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	cryptokeyring "github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdktx "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// KeyPair defines a public/private key pair to e.g. sign a transaction.
type KeyPair struct {
	Pub  cryptotypes.PubKey
	Priv cryptotypes.PrivKey
}

// Result defines a generic way to receive responses from the RPC endpoint.
type Result struct {
	Info map[string]interface{} `mapstructure:"info" json:"info"`
}

func init() {
	GetConfig()
}

func getKeyPairFromKeyring(address sdk.AccAddress) (keyPair KeyPair, err error) {
	ring, err := keyring.Open(keyring.Config{
		AllowedBackends: []keyring.BackendType{keyring.FileBackend},
		FileDir:         filepath.Join(libConfig.RootDir, "keyring-test"),
		FilePasswordFunc: func(_ string) (string, error) {
			return "test", nil
		},
	})
	if err != nil {
		return
	}

	name := hex.EncodeToString([]byte(address)) + ".address"
	i, err := ring.Get(name)
	if err != nil {
		return
	}

	name = string(i.Data)
	i, err = ring.Get(name)
	if err != nil {
		return
	}

	s := hex.EncodeToString(i.Data)
	privKey := s[len(s)-64:]

	decodedPriv, err := hex.DecodeString(privKey)
	if err != nil {
		return
	}

	algo, err := cryptokeyring.NewSigningAlgoFromString("secp256k1", cryptokeyring.SigningAlgoList{hd.Secp256k1})
	if err != nil {
		return
	}

	priv := algo.Generate()(decodedPriv)
	pub := priv.PubKey()

	keyPair = KeyPair{
		Pub:  pub,
		Priv: priv,
	}
	return
}

func getAccountNumberAndSequence(address sdk.AccAddress) (accountNumber, sequence uint64, err error) {
	grpcConn, err := libConfig.GetGRPCConn()
	if err != nil {
		return
	}
	defer grpcConn.Close()

	client := authtypes.NewQueryClient(grpcConn)
	grpcRes, err := client.AccountInfo(
		context.Background(),
		&authtypes.QueryAccountInfoRequest{
			Address: address.String(),
		},
	)
	if err != nil {
		return
	}

	accountNumber = grpcRes.Info.AccountNumber
	sequence = grpcRes.Info.Sequence
	return
}

// BuildAndSignTx constructs the transaction from address' private key and messages.
func BuildAndSignTx(address sdk.AccAddress, msgs ...sdk.Msg) (txBytes []byte, txJSON string, err error) {
	encodingConfig := GetConfig().EncodingConfig
	if encodingConfig.TxConfig == nil {
		err = errors.New("encoding config must not be nil")
		return
	}
	txBuilder := encodingConfig.TxConfig.NewTxBuilder()

	err = txBuilder.SetMsgs(msgs...)
	if err != nil {
		return
	}

	txBuilder.SetFeeAmount(sdk.Coins{sdk.NewInt64Coin("plmnt", 1)})
	txBuilder.SetGasLimit(200000)
	txBuilder.SetTimeoutHeight(0)

	keyPair, err := getKeyPairFromKeyring(address)
	if err != nil {
		return
	}

	accountNumber, sequence, err := getAccountNumberAndSequence(address)
	if err != nil {
		return
	}

	// First round: we gather all the signer infos. We use the "set empty signature" hack to do that.
	var sigsV2 []signing.SignatureV2
	sigV2 := signing.SignatureV2{
		PubKey: keyPair.Pub,
		Data: &signing.SingleSignatureData{
			SignMode:  encodingConfig.TxConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: sequence,
	}
	sigsV2 = append(sigsV2, sigV2)
	err = txBuilder.SetSignatures(sigsV2...)
	if err != nil {
		return
	}

	// Second round: all signer infos are set, so each signer can sign.
	sigsV2 = []signing.SignatureV2{}
	signerData := xauthsigning.SignerData{
		ChainID:       libConfig.ChainID,
		AccountNumber: accountNumber,
		Sequence:      sequence,
	}
	sigV2, err = tx.SignWithPrivKey(encodingConfig.TxConfig.SignModeHandler().DefaultMode(), signerData, txBuilder, keyPair.Priv, encodingConfig.TxConfig, sequence)
	if err != nil {
		return
	}
	sigsV2 = append(sigsV2, sigV2)
	err = txBuilder.SetSignatures(sigsV2...)
	if err != nil {
		return
	}

	// Generated Protobuf-encoded bytes.
	txBytes, err = encodingConfig.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return
	}

	// Generate a JSON string.
	txJSONBytes, err := encodingConfig.TxConfig.TxJSONEncoder()(txBuilder.GetTx())
	if err != nil {
		return
	}
	txJSON = string(txJSONBytes)
	return
}

// BroadcastTx broadcasts a transaction via gRPC.
func BroadcastTx(txBytes []byte) (txResponse *sdk.TxResponse, err error) {
	grpcConn, err := libConfig.GetGRPCConn()
	if err != nil {
		return
	}
	defer grpcConn.Close()

	client := sdktx.NewServiceClient(grpcConn)
	grpcRes, err := client.BroadcastTx(
		context.Background(),
		&sdktx.BroadcastTxRequest{
			Mode:    sdktx.BroadcastMode_BROADCAST_MODE_SYNC,
			TxBytes: txBytes,
		},
	)
	if err != nil {
		return
	}
	txResponse = grpcRes.TxResponse
	return
}

// SimulateTx simulates broadcasting a transaction via gRPC.
func SimulateTx(txBytes []byte) (result *sdk.Result, err error) {
	grpcConn, err := libConfig.GetGRPCConn()
	if err != nil {
		return
	}
	defer grpcConn.Close()

	client := sdktx.NewServiceClient(grpcConn)
	grpcRes, err := client.Simulate(
		context.Background(),
		&sdktx.SimulateRequest{
			TxBytes: txBytes,
		},
	)
	if err != nil {
		return
	}
	result = grpcRes.Result
	return
}
