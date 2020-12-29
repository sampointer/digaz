package command

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const ipv4 = "13.66.60.119"
const ipv6 = "2603:1000:4:402::179"

func TestLookup(t *testing.T) {
	t.Parallel()
	t.Run("looks up IPv4 address", func(t *testing.T) {
		t.Parallel()

		p, err := Lookup(ipv4)
		require.NoError(t, err)
		require.Equal(t, 4, len(p))
	})

	t.Run("looks up IPv6 address", func(t *testing.T) {
		t.Parallel()

		p, err := Lookup(ipv6)
		require.NoError(t, err)
		require.Equal(t, 4, len(p))
	})
}
