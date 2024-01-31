package e2e

import (
	"errors"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/planetmint/planetmint-go/lib"
	clitestutil "github.com/planetmint/planetmint-go/testutil/cli"
	"github.com/planetmint/planetmint-go/testutil/network"
	"github.com/planetmint/planetmint-go/testutil/sample"
	machinetypes "github.com/planetmint/planetmint-go/x/machine/types"
)

func CreateAccount(network *network.Network, name string, mnemonic string) (account *keyring.Record, err error) {
	val := network.Validators[0]

	kb := val.ClientCtx.Keyring
	account, err = kb.NewAccount(name, mnemonic, keyring.DefaultBIP39Passphrase, sample.DefaultDerivationPath, hd.Secp256k1)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func FundAccount(network *network.Network, account *keyring.Record) (err error) {
	val := network.Validators[0]

	addr, err := account.GetAddress()
	if err != nil {
		return err
	}

	// sending funds to account to initialize account on chain
	coin := sdk.NewCoins(sdk.NewInt64Coin("stake", 10000)) // TODO: make denom dependent on cfg
	msg := banktypes.NewMsgSend(val.Address, addr, coin)
	out, err := lib.BroadcastTxWithFileLock(val.Address, msg)
	if err != nil {
		return err
	}

	err = network.WaitForNextBlock()
	if err != nil {
		return err
	}

	rawLog, err := clitestutil.GetRawLogFromTxOut(val, out)
	if err != nil {
		return err
	}

	if !strings.Contains(rawLog, "cosmos.bank.v1beta1.MsgSend") {
		err = errors.New("failed to fund account")
	}

	return
}

func AttestMachine(network *network.Network, name string, mnemonic string, num int) (err error) {
	val := network.Validators[0]

	account, err := CreateAccount(network, name, mnemonic)
	if err != nil {
		return err
	}

	err = FundAccount(network, account)
	if err != nil {
		return err
	}

	// register Ta
	prvKey, pubKey := sample.KeyPair(num)

	ta := sample.TrustAnchor(pubKey)
	registerMsg := machinetypes.NewMsgRegisterTrustAnchor(val.Address.String(), &ta)
	_, err = lib.BroadcastTxWithFileLock(val.Address, registerMsg)
	if err != nil {
		return err
	}
	err = network.WaitForNextBlock()
	if err != nil {
		return err
	}

	addr, err := account.GetAddress()
	if err != nil {
		return err
	}

	machine := sample.Machine(name, pubKey, prvKey, addr.String())
	attestMsg := machinetypes.NewMsgAttestMachine(addr.String(), &machine)
	_, err = lib.BroadcastTxWithFileLock(addr, attestMsg)
	if err != nil {
		return err
	}

	err = network.WaitForNextBlock()
	if err != nil {
		return err
	}

	return
}
