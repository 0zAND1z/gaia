package main

import (
	"os"

	dg "github.com/cosmos/basecoin-delegationgame"
	"github.com/spf13/cobra"
	// _ "github.com/tendermint/basecoin-delegationgame/commands"
	_ "github.com/tendermint/basecoin-stake/cmd/stakecoin/commands"
	"github.com/tendermint/basecoin/cmd/commands"
	"github.com/tendermint/basecoin/types"
	"github.com/tendermint/tmlibs/cli"
)

func init() {
	commands.RegisterStartPlugin("delegationGame",
		func() types.Plugin {
			return &dg.Plugin{}
		},
	)
}

func main() {
	var RootCmd = &cobra.Command{
		Use: "gaia",
	}

	RootCmd.AddCommand(
		commands.InitCmd,
		commands.StartCmd,
		commands.TxCmd,
		commands.QueryCmd,
		commands.KeyCmd,
		commands.VerifyCmd,
		commands.BlockCmd,
		commands.AccountCmd,
		commands.UnsafeResetAllCmd,
		commands.QuickVersionCmd("0.1.0"),
	)

	cli.PrepareMainCmd(
		RootCmd,
		"GAIA", // envvar prefix
		os.ExpandEnv("$HOME/.cosmos-gaia"), // default home
	).Execute()
}