package ipforward

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thxcode/winnet/pkg/adapter"
)

func TestGetNetRoutesAll(t *testing.T) {
	routes, err := GetNetRoutesAll()
	assert.NoError(t, err)

	t.Logf("%4s\t %18s\t %18s\t %6s\n", "Iface", "Dest", "Gateway", "Metric")
	for _, route := range routes {
		t.Logf("%4d\t %18s\t %18s\t %6d\n", route.LinkIndex, route.DestinationSubnet, route.GatewayAddress, route.RouteMetric)
	}
}

func TestNewGetDeleteNetRoute(t *testing.T) {
	defaultIfaceIdx, err := adapter.GetDefaultGatewayIfaceIndex()
	assert.NoError(t, err)

	tests := []struct {
		giveLinkIndex int
		giveDstSubnet string
		giveGateway   string
	}{
		{
			giveLinkIndex: defaultIfaceIdx,
			giveDstSubnet: "7.7.7.7/32",
			giveGateway:   "1.2.3.4",
		},
		{
			giveLinkIndex: defaultIfaceIdx,
			giveDstSubnet: "8.7.8.7/32",
			giveGateway:   "8.8.8.8",
		},
	}

	for _, tt := range tests {
		_, dstSubnet, err := net.ParseCIDR(tt.giveDstSubnet)
		assert.NoError(t, err)

		gateway := net.ParseIP(tt.giveGateway)

		err = NewNetRoute(tt.giveLinkIndex, dstSubnet, gateway)
		assert.NoError(t, err)

		routes, err := GetNetRoutes(tt.giveLinkIndex, dstSubnet)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(routes))
		assert.Equal(t, tt.giveLinkIndex, routes[0].LinkIndex)
		assert.Equal(t, tt.giveDstSubnet, routes[0].DestinationSubnet.String())
		assert.Equal(t, tt.giveGateway, routes[0].GatewayAddress.String())
		assert.Less(t, DefaultRouteMetric, routes[0].RouteMetric)

		err = RemoveNetRoute(tt.giveLinkIndex, dstSubnet, gateway)
		assert.NoError(t, err)

		routes, err = GetNetRoutes(tt.giveLinkIndex, dstSubnet)
		assert.NoError(t, err)
		assert.Empty(t, routes)
	}
}
