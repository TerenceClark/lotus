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
	Flags: []cli.Flag{
		&cli.StringSliceFlag{Name: "hosts"},
	},
	Action: func(cctx *cli.Context) error {
		api, closer, err := GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		return api.SetValidHosts(ReqContext(cctx), cctx.StringSlice("hosts"))
	},
}
