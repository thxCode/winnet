package nets

import (
	"fmt"
	"github.com/thxcode/winnet/cmd/utils"
	"github.com/thxcode/winnet/pkg/adapter"
	"github.com/urfave/cli"
)

var getFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "def-gw",
		Usage: "[optional] Get default gateway name",
	},
	cli.StringFlag{
		Name:  "name",
		Usage: "[optional] Get interface by name",
	},
	cli.IntFlag{
		Name:  "idx",
		Usage: "[optional] Get interface by index",
	},
	cli.StringFlag{
		Name:  "ip",
		Usage: "[optional] Get interface by ip",
	},
}

func getAction(cliCtx *cli.Context) (err error) {
	getDefGw := cliCtx.Bool("def-gw")
	byName := cliCtx.String("name")
	byIndex := cliCtx.Int("idx")
	byIP := cliCtx.String("ip")
	stdw := cliCtx.App.Writer

	if getDefGw {
		ifaceName, err := adapter.GetDefaultGatewayIfaceName()
		if err != nil {
			return err
		}
		return utils.Output(stdw, ifaceName)
	}
	if byName != "" {
		iface, err := adapter.GetInterfaceByName(byName)
		if err != nil {
			return err
		}
		return utils.Output(stdw, iface)
	}
	if byIndex != 0 {
		if byIndex < 0 {
			return fmt.Errorf("the index of interface should not be nagitive: %d", byIndex)
		}

		iface, err := adapter.GetInterfaceByIndex(byIndex)
		if err != nil {
			return err
		}
		return utils.Output(stdw, iface)
	}
	if byIP != "" {
		iface, err := adapter.GetInterfaceByIP(byIP)
		if err != nil {
			return err
		}
		return utils.Output(stdw, iface)
	}

	// all
	ifaces, err := adapter.GetInterfaces()
	if err != nil {
		return err
	}
	return utils.Output(stdw, ifaces)
}

func get() cli.Command {
	return cli.Command{
		Name:   "get",
		Flags:  getFlags,
		Action: getAction,
	}
}
