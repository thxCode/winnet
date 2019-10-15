package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInterfaces(t *testing.T) {
	ifaces, err := GetInterfaces()
	assert.NoError(t, err)

	t.Logf("%4s\t %20s\t %40s\t %15s\t %15s\t %15s\t %18s\t %5s\n", "Index", "Name", "Description", "Addr", "Mask", "Gateway", "SubnetCIDR", "DCHP")
	for _, iface := range ifaces {
		t.Logf("%4d\t %20s\t %40s\t %15s\t %15s\t %15s\t %18s\t %5v\n",
			iface.Idx,
			iface.Name,
			iface.Description,
			iface.IpAddress,
			iface.IpMask,
			iface.DefaultGatewayAddress,
			iface.GetSubnet(),
			iface.DhcpEnabled,
		)
	}
}

func TestGetInterface(t *testing.T) {
	// get default name
	defaultIfaceName, err := GetDefaultGatewayIfaceName()
	assert.NoError(t, err)
	assert.NotEmpty(t, defaultIfaceName)

	// get default index
	defaultIfaceIdx, err := GetDefaultGatewayIfaceIndex()
	assert.NoError(t, err)
	assert.NotZero(t, defaultIfaceIdx)

	// get by name
	defaultIface, err := GetInterfaceByName(defaultIfaceName)
	assert.NoError(t, err)
	assert.NotNil(t, defaultIface)
	assert.Equal(t, defaultIfaceIdx, defaultIface.Idx)

	// get all
	ifaces, err := GetInterfaces()
	assert.NoError(t, err)
	assert.NotEmpty(t, ifaces)
	assert.Contains(t, ifaces, defaultIface)

	// get by idx
	iface, err := GetInterfaceByIndex(defaultIface.Idx)
	assert.NoError(t, err)
	assert.NotNil(t, iface)
	assert.Equal(t, defaultIface.IpAddress, iface.IpAddress)
	assert.Equal(t, defaultIface.Name, iface.Name)

	// get by ip
	iface, err = GetInterfaceByIP(defaultIface.IpAddress)
	assert.NoError(t, err)
	assert.NotNil(t, iface)
	assert.Equal(t, defaultIface.Idx, iface.Idx)
	assert.Equal(t, defaultIface.Name, iface.Name)
}

func TestEnableForwarding(t *testing.T) {
	// get default name
	defaultIfaceName, err := GetDefaultGatewayIfaceName()
	assert.NoError(t, err)
	assert.NotEmpty(t, defaultIfaceName)

	err = EnableForwardingByName(defaultIfaceName)
	assert.NoError(t, err)

	// get default index
	defaultIfaceIdx, err := GetDefaultGatewayIfaceIndex()
	assert.NoError(t, err)
	assert.NotZero(t, defaultIfaceIdx)

	err = EnableForwardingByIndex(defaultIfaceIdx)
	assert.NoError(t, err)
}
