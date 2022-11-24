package gophercon2021

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountryDB(t *testing.T) {
	w := Stem("threading")
	require.Equal(t, "thread", w)
}
