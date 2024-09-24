package cmd

import (
	"fmt"
	"strings"

	"github.com/cosmos/go-bip39"
	"github.com/planetmint/planetmint-go/lib/trustwallet"
	"github.com/spf13/cobra"
)

func InitializeTrustWalletCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "initialize-trust-wallet [mnemonic_word_1][,[mnemonic_word_2],...[mnemonic_word_12],...[mnemonic_word_24]]",
		Short: "Initialize a trust wallet",
		Long:  "TODO: add long description",
		RunE:  initializeTrustWalletCmdFunc,
		Args:  cobra.RangeArgs(0, 1),
	}

	return cmd
}

func initializeTrustWalletCmdFunc(cmd *cobra.Command, args []string) error {
	// TODO: read portName from config
	connector, err := trustwallet.NewTrustWalletConnector("/dev/ttyACM0")
	if err != nil {
		return err
	}

	// if no mnemonic is given create one
	if len(args) == 0 {
		fmt.Println("initializing Trust Wallet")
		mnemonic, err := connector.CreateMnemonic()
		if err != nil {
			return err
		}
		fmt.Println(mnemonic)

		return nil
	}

	fmt.Println("recovering Trust Wallet from mnemonic...")
	words := strings.Split(args[0], ",")
	if len(words) != 12 && len(words) != 24 {
		return fmt.Errorf("expected length of mnemonic is 12 or 24, got: %d", len(words))
	}
	mnemonic := strings.Join(words, " ")
	if !bip39.IsMnemonicValid(mnemonic) {
		return fmt.Errorf("invalid mnemonic, please check provided words")
	}
	response, err := connector.RecoverFromMnemonic(mnemonic)
	if err != nil {
		return err
	}
	fmt.Println(response)

	return nil
}
