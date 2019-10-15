package routes

import (
	"errors"
	"fmt"
	gonet "net"

	"github.com/thxcode/winnet/cmd/utils"
	"github.com/thxcode/winnet/pkg/ipforward"
	"github.com/urfave/cli"
)

var getFlags = []cli.Flag{
	cli.IntFlag{
		Name:  "iface",
		Usage: "[optional] Indicate the index of interface",
	},
	cli.StringFlag{
		Name:  "dst",
		Usage: "[optional] Indicate the CIDR of destination",
	},
}

func getAction(cliCtx *cli.Context) (err error) {
	ifaceIdx := cliCtx.Int("iface")
	routeDest := cliCtx.String("dst")
	stdw := cliCtx.App.Writer

	var routes []ipforward.Route

	if ifaceIdx != 0 || routeDest != "" {
		if ifaceIdx < 0 {
			return fmt.Errorf("the index of interface should not be nagitive: %d", ifaceIdx)
		}

		if routeDest == "" || ifaceIdx == 0 {
			return errors.New("indicate both 'iface' and 'dst'")
		}

		_, routeDestIpNet, err := gonet.ParseCIDR(routeDest)
		if err != nil {
			return fmt.Errorf("destination %s is an invaild CIDR format", routeDest)
		}
		routes, err = ipforward.GetNetRoutes(ifaceIdx, routeDestIpNet)
	} else {
		routes, err = ipforward.GetNetRoutesAll()
	}

	// all
	if err != nil {
		return err
	}
	return utils.Output(stdw, routes)
}

func get() cli.Command {
	return cli.Command{
		Name:   "get",
		Flags:  getFlags,
		Action: getAction,
	}
}
