package service_tags

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

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
	})
}
