package util

import (
	"log"
	"strconv"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
	"github.com/planetmint/planetmint-go/config"
)

// TODO: define the final issuance
func GetUnsignedReissuanceTransaction() (string, error) {
	config := config.GetConfig()
	// Create a new client instance

	connCfg := &rpcclient.ConnConfig{
		Host:         config.RPCHost + ":" + strconv.Itoa(config.RPCPort),
		User:         config.RPCUser,
		Pass:         config.RPCPassword,
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}

	client, erro := rpcclient.New(connCfg, nil)
	if erro != nil {
		log.Fatal(erro)
	}
	defer client.Shutdown()

	// Define the inputs and outputs for the raw transaction
	txIn := wire.NewTxIn(&wire.OutPoint{
		Hash:  chainhash.Hash{},
		Index: 0,
	}, nil, nil)

	txOut := wire.NewTxOut(1000, nil) // 1000 satoshis, replace with your desired amount

	// Create the raw transaction
	rawTx := wire.NewMsgTx(wire.TxVersion)
	rawTx.AddTxIn(txIn)
	rawTx.AddTxOut(txOut)

	// Serialize the transaction and convert to hex
	//buf := []byte{}
	var txHex = "12341234123412312342342341234123412341231234234234"
	return txHex, erro
}
