package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/go-bip39"
	"github.com/planetmint/planetmint-go/lib/trustwallet"
	"github.com/spf13/cobra"
)

const (
	flagSerialPort = "serial-port"
)

func InitializeTrustWalletCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "initialize-trust-wallet [mnemonic_word_1][,[mnemonic_word_2],...[mnemonic_word_12],...[mnemonic_word_24]]",
		Short: "Initialize a Trust Wallet",
		Long: `Initialize a Trust Wallet with a mnemonic phrase (optional). If no mnemonic is provided then one is created for you. 
Provided mnemonics must be 12 or 24 words long and adhere to bip39.`,
		RunE: initializeTrustWalletCmdFunc,
		Args: cobra.RangeArgs(0, 1),
	}

	cmd.Flags().String(flagSerialPort, "/dev/ttyACM0", "The serial port your Trust Wallet is connected to")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func initializeTrustWalletCmdFunc(cmd *cobra.Command, args []string) error {
	serialPort, err := cmd.Flags().GetString(flagSerialPort)
	if err != nil {
		return err
	}

	connector, err := trustwallet.NewTrustWalletConnector(serialPort)
	if err != nil {
		return err
	}

	// create mnemonic if non is given
	if len(args) == 0 {
		fmt.Println("Initializing Trust Wallet. This may take a few seconds...")
		mnemonic, err := connector.CreateMnemonic()
		if err != nil {
			return err
		}

		fmt.Println("Created mnemonic:")
		fmt.Println(mnemonic + "\n")
		fmt.Println("IMPORTANT: Store your mnemonic securely in an offline location, such as a hardware wallet, encrypted USB, or paper stored in a safe, never online or on cloud storage!")

		return nil
	}

	// recover from given mnemonic
	fmt.Println("Recovering Trust Wallet from mnemonic...")
	words := strings.Split(args[0], ",")
	if len(words) != 12 && len(words) != 24 {
		return errors.New("expected length of mnemonic is 12 or 24, got: " + strconv.Itoa(len(words)))
	}
	mnemonic := strings.Join(words, " ")
	if !bip39.IsMnemonicValid(mnemonic) {
		return errors.New("invalid mnemonic, please check provided words")
	}
	response, err := connector.RecoverFromMnemonic(mnemonic)
	if err != nil {
		return err
	}
	fmt.Println(response)

	return nil
}
