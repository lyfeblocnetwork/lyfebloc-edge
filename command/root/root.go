package root

import (
	"fmt"
	"os"

	"github.com/lyfeblocnetwork/lyfebloc-edge/command/backup"
	"github.com/lyfeblocnetwork/lyfebloc-edge/command/genesis"
	"github.com/lyfeblocnetwork/lyfebloc-edge/command/helper"
	"github.com/lyfeblocnetwork/lyfebloc-edge/command/ibft"
	"github.com/lyfeblocnetwork/lyfebloc-edge/command/license"
	"github.com/lyfeblocnetwork/lyfebloc-edge/command/monitor"
	"github.com/lyfeblocnetwork/lyfebloc-edge/command/peers"
	"github.com/lyfeblocnetwork/lyfebloc-edge/command/secrets"
	"github.com/lyfeblocnetwork/lyfebloc-edge/command/server"
	"github.com/lyfeblocnetwork/lyfebloc-edge/command/status"
	"github.com/lyfeblocnetwork/lyfebloc-edge/command/txpool"
	"github.com/lyfeblocnetwork/lyfebloc-edge/command/version"
	"github.com/lyfeblocnetwork/lyfebloc-edge/command/whitelist"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Lyfebloc Edge is a framework for building Ethereum-compatible Blockchain networks",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		monitor.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		whitelist.GetCommand(),
		license.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
