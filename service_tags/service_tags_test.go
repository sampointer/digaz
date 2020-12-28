package service_tags

import (
	"encoding/json"
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
		res, props, err := st.LookupIPv4(net.ParseIP("13.66.60.119"))
		require.NoError(t, err)
		require.Equal(t, 4, len(res))
		require.Equal(t, res[0].Name, "ActionGroup")
		require.Equal(t, res[0].Id, "ActionGroup")

		require.Equal(t, 4, len(props))

		for _, x := range props {
			t.Log(x.String())
		}
	})

	t.Run("does not look up IPv6", func(t *testing.T) {
		_, _, err := st.LookupIPv4(net.ParseIP("2603:1000:4:402::179"))
		require.Error(t, err)
	})
}

func TestLookupIPv6(t *testing.T) {
	t.Parallel()

	doc, err := os.Open("../testdata/ServiceTags_Public_20201214.json")
	require.NoError(t, err)

	st, err := New(doc)
	require.NoError(t, err)

	t.Run("looks up valid IPv6", func(t *testing.T) {
		res, props, err := st.LookupIPv6(net.ParseIP("2603:1000:4:402::179"))
		require.NoError(t, err)
		require.Equal(t, 4, len(res))
		require.Equal(t, res[0].Name, "ActionGroup")
		require.Equal(t, res[0].Id, "ActionGroup")

		require.Equal(t, 4, len(props))
	})

	t.Run("does not look up IPv4", func(t *testing.T) {
		_, _, err := st.LookupIPv6(net.ParseIP("13.66.60.119"))
		require.Error(t, err)
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

			t.Run("is a Stringer", func(t *testing.T) {
				t.Parallel()
				require.Equal(t,
					"changeNumber: 5 networkFeatures: [\"API\" \"NSG\" \"UDR\" \"FW\"] platform: \"Azure\" region: \"\" regionId: 0 systemService: \"ActionGroup\"",
					st.Values[0].Properties.String(),
				)
			})

			t.Run("implements JSON()", func(t *testing.T) {
				t.Parallel()
				e, err := json.Marshal(st.Values[0].Properties)
				require.NoError(t, err)

				j, err := st.Values[0].Properties.JSON()
				require.NoError(t, err)
				require.Equal(t, string(e), j)
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
