package routes

import (
	"github.com/urfave/cli"
)

func NewCommand() cli.Command {
	return cli.Command{
		Name: "routes",
		Subcommands: []cli.Command{
			get(),
			create(),
			delete(),
		},
	}
}
