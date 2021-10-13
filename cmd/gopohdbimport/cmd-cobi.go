package main

import (
	"strings"

	"github.com/mitchellh/cli"
)

type cobiCmd struct {
	Ui cli.Ui
}

func (cmd *cobiCmd) Help() string {
	helpText := `
Usage: gopohdbimport cobi nama_tabel [options]
	Dispatches a custom event across the Serf cluster.
Options:
	-d=dbEnvKey             (default "DEFAULT")
`
	return strings.TrimSpace(helpText)
}

func (cmd *cobiCmd) Run(args []string) int {



	return 0
}

func (cmd *cobiCmd) Synopsis() string {
	return "Send a custom event through the Serf cluster"
}

func init() {
	cmds["cobi"] = func() (cli.Command, error) {
		return &cobiCmd{Ui: ui}, nil
	}
}
