package nets

import (
	"github.com/urfave/cli"
)

func NewCommand() cli.Command {
	return cli.Command{
		Name: "nets",
		Subcommands: []cli.Command{
			get(),
			forward(),
		},
	}
}
