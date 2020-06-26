package cli

import (
	"github.com/urfave/cli/v2"
)

var xCtrlCmd = &cli.Command{
	Name:  "xctrl",
	Usage: "Customized commands",
	Subcommands: []*cli.Command{
		setValidHosts,
	},
}

var setValidHosts = &cli.Command{
	Name:  "setvalidhosts",
	Usage: "Set valid hosts",
	Action: func(cctx *cli.Context) error {

		return nil
	},
}
