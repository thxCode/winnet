package routes

import (
	"errors"
	"fmt"
	gonet "net"

	"github.com/thxcode/winnet/pkg/ipforward"
	"github.com/urfave/cli"
)

var createFlags = []cli.Flag{
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
	createIfaceIdx     int
	createRouteDest    string
	createRouteGateway string
)

func createParser(cliCtx *cli.Context) error {
	createIfaceIdx = cliCtx.Int("iface")
	createRouteDest = cliCtx.String("dst")
	createRouteGateway = cliCtx.String("gw")

	if createIfaceIdx == 0 || createRouteDest == "" || createRouteGateway == "" {
		return errors.New("must indicate all of 'iface', 'dst', 'gw'")
	}

	return nil
}

func createAction(cliCtx *cli.Context) (err error) {
	_, routeDestIpNet, err := gonet.ParseCIDR(createRouteDest)
	if err != nil {
		return fmt.Errorf("destination %s is an invaild CIDR format", createRouteDest)
	}
	routeGatewayIp := gonet.ParseIP(createRouteGateway)

	return ipforward.NewNetRoute(createIfaceIdx, routeDestIpNet, routeGatewayIp)
}

func create() cli.Command {
	return cli.Command{
		Name:   "create",
		Flags:  createFlags,
		Before: createParser,
		Action: createAction,
	}
}
