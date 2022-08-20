package flag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAny(t *testing.T) {
	got := Any("/tmp")

	require.NotNil(t, got)
}
