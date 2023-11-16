package main

import (
	"errors"
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/planetmint/planetmint-go/app"
	"github.com/planetmint/planetmint-go/cmd/planetmint-god/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		var e *server.ErrorCode
		if errors.As(err, &e) {
			os.Exit(e.Code)
		}
		os.Exit(1)
	}
}
