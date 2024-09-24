package cmd

import (
	"errors"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/go-bip39"
	"github.com/planetmint/planetmint-go/lib/trustwallet"
	"github.com/spf13/cobra"
)

const (
	flagSerialPort = "serial-port"
)

func TrustWalletCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "trust-wallet [command]",
		Short:                      "Trust Wallet subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		initializeCmd(),
		keysCmd(),
	)

	return cmd
}

func initializeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "initialize [mnemonic_word_1][,[mnemonic_word_2],...[mnemonic_word_12],...[mnemonic_word_24]]",
		Short: "Initialize a Trust Wallet",
		Long: `Initialize a Trust Wallet with a mnemonic phrase (optional). If no mnemonic is provided then one is created for you. 
Provided mnemonics must be 12 or 24 words long and adhere to bip39.`,
		RunE: initializeCmdFunc,
		Args: cobra.RangeArgs(0, 1),
	}

	cmd.Flags().String(flagSerialPort, "/dev/ttyACM0", "The serial port your Trust Wallet is connected to")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func initializeCmdFunc(cmd *cobra.Command, args []string) error {
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
		cmd.Println("Initializing Trust Wallet. This may take a few seconds...")
		mnemonic, err := connector.CreateMnemonic()
		if err != nil {
			return err
		}

		cmd.Println("Created mnemonic:")
		cmd.Println(mnemonic + "\n")
		cmd.Println("IMPORTANT: Store your mnemonic securely in an offline location, such as a hardware wallet, encrypted USB, or paper stored in a safe, never online or on cloud storage!")

		return nil
	}

	// recover from given mnemonic
	cmd.Println("Recovering Trust Wallet from mnemonic...")
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
	cmd.Println(response)

	return nil
}

func keysCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keys",
		Short: "Retrieve keys from Trust Wallet",
		Long: `Retrieve keys from Trust Wallet. Includes: 
planetmint address,
extended planetmint public key,
extended liquid public key,
raw planetmint key (hex encoded)`,
		RunE: keysCmdFunc,
		Args: cobra.ExactArgs(0),
	}

	cmd.Flags().String(flagSerialPort, "/dev/ttyACM0", "The serial port your Trust Wallet is connected to")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func keysCmdFunc(cmd *cobra.Command, _ []string) error {
	serialPort, err := cmd.Flags().GetString(flagSerialPort)
	if err != nil {
		return err
	}

	connector, err := trustwallet.NewTrustWalletConnector(serialPort)
	if err != nil {
		return err
	}

	cmd.Println("Retrieving keys from Trust Wallet. This may take a few seconds...")
	cmd.Println()

	keys, err := connector.GetPlanetmintKeys()
	if err != nil {
		return err
	}

	cmd.Println("Planetmint address:      " + keys.PlanetmintAddress)
	cmd.Println("Planetmint public key:   " + keys.ExtendedPlanetmintPubkey)
	cmd.Println("Liquid public key:       " + keys.ExtendedLiquidPubkey)
	cmd.Println("Raw Planetmint key(hex): " + keys.RawPlanetmintPubkey)

	return nil
}
