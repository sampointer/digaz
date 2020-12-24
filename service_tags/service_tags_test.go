package service_tags

import (
	"net"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLookupIPv4(t *testing.T) {
	t.Parallel()

	doc, err := os.Open("../testdata/ServiceTags_Public_20201214.json")
	require.NoError(t, err)

	st, err := New(doc)
	require.NoError(t, err)

	t.Run("looks up valid IPv4", func(t *testing.T) {
		res, err := st.LookupIPv4(net.ParseIP("13.66.60.119"))
		require.NoError(t, err)
		require.NotZero(t, len(res))
		// TODO: assert on result object equality with fixture
	})

}

func TestServiceTags(t *testing.T) {
	t.Parallel()

	doc, err := os.Open("../testdata/ServiceTags_Public_20201214.json")
	require.NoError(t, err)

	st, err := New(doc)
	require.NoError(t, err)

	t.Run("has change number", func(t *testing.T) {
		t.Parallel()
		require.NotZero(t, st.ChangeNumber, "should have a non-zero change number")
	})

	t.Run("has a cloud", func(t *testing.T) {
		t.Parallel()
		require.NotZero(t, st.Cloud, "should have a cloud")
	})

	t.Run("has values", func(t *testing.T) {
		t.Parallel()
		require.NotZero(t, len(st.Values), "should have a least one value")

		t.Run("value has name and ID", func(t *testing.T) {
			t.Parallel()
			require.NotZero(t, st.Values[0].Name, "has a name")
			require.NotZero(t, st.Values[0].Id, "has an ID")
		})

		t.Run("has properties", func(t *testing.T) {
			t.Parallel()
			require.NotZero(t, st.Values[0].Properties, "has properties")

			t.Run("has network features", func(t *testing.T) {
				t.Parallel()
				require.NotZero(t, len(st.Values[0].Properties.NetworkFeatures))
			})

			t.Run("has flat properties", func(t *testing.T) {
				t.Parallel()
				require.NotZero(t, st.Values[0].Properties.ChangeNumber)
				require.NotZero(t, st.Values[0].Properties.Platform)
				require.NotZero(t, st.Values[0].Properties.SystemService)
			})

			t.Run("has address prefixes", func(t *testing.T) {
				t.Parallel()
				require.NotZero(t, len(st.Values[0].Properties.AddressPrefixes))
			})

		})
	})
}
