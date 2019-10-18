package tool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteNetsh(t *testing.T) {
	// input an empty argument list
	result, _, err := ExecuteNetsh(nil)
	assert.Error(t, err)
	assert.False(t, result)

	// netsh int ipv4 show addresses
	result, _, err = ExecuteNetsh([]string{"int", "ipv4", "show", "addresses"})
	assert.NoError(t, err)
	assert.True(t, result)
}
