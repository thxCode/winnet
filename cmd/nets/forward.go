package nets

import (
	"errors"

	"github.com/thxcode/winnet/pkg/adapter"
	"github.com/urfave/cli"
)

var forwardFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "name",
		Usage: "[optional] Enable the forwarding of the given name interface",
	},
	cli.IntFlag{
		Name:  "idx",
		Usage: "[optional] Enable the forwarding of the given index interface",
	},
}

var (
	byName  string
	byIndex int
)

func forwardParser(cliCtx *cli.Context) error {
	byName = cliCtx.String("name")
	byIndex = cliCtx.Int("idx")

	if byName == "" && byIndex == 0 {
		return errors.New("could not enable the forwarding of an unknown interface")
	}

	return nil
}

func forwardAction(cliCtx *cli.Context) (err error) {
	if byName != "" {
		return adapter.EnableForwardingByName(byName)
	}
	if byIndex != 0 {
		return adapter.EnableForwardingByIndex(byIndex)
	}

	return nil
}

func forward() cli.Command {
	return cli.Command{
		Name:   "forward",
		Flags:  forwardFlags,
		Before: forwardParser,
		Action: forwardAction,
	}
}
