package routes

import (
	"errors"
	"fmt"
	gonet "net"

	"github.com/thxcode/winnet/pkg/ipforward"
	"github.com/urfave/cli"
)

var deleteFlags = []cli.Flag{
	cli.IntFlag{
		Name:  "iface",
		Usage: "[required] Indicate the index of interface",
	},
	cli.StringFlag{
		Name:  "dst",
		Usage: "[required] Indicate the CIDR of destination",
	},
	cli.StringFlag{
		Name:  "gw",
		Usage: "[required] Indicate the address of gateway",
	},
}

var (
	deleteIfaceIdx     int
	deleteRouteDest    string
	deleteRouteGateway string
)

func deleteParser(cliCtx *cli.Context) error {
	deleteIfaceIdx = cliCtx.Int("iface")
	deleteRouteDest = cliCtx.String("dst")
	deleteRouteGateway = cliCtx.String("gw")

	if deleteIfaceIdx == 0 || deleteRouteDest == "" || deleteRouteGateway == "" {
		return errors.New("must indicate all of 'iface', 'dst', 'gw'")
	}

	return nil
}

func deleteAction(cliCtx *cli.Context) (err error) {
	_, routeDestIpNet, err := gonet.ParseCIDR(deleteRouteDest)
	if err != nil {
		return fmt.Errorf("destination %s is an invaild CIDR format", deleteRouteDest)
	}
	routeGatewayIp := gonet.ParseIP(deleteRouteGateway)

	return ipforward.RemoveNetRoute(deleteIfaceIdx, routeDestIpNet, routeGatewayIp)
}

func delete() cli.Command {
	return cli.Command{
		Name:   "delete",
		Flags:  deleteFlags,
		Before: deleteParser,
		Action: deleteAction,
	}
}
