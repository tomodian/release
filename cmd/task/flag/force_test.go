package flag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestForce(t *testing.T) {
	got := Force("/tmp")

	require.NotNil(t, got)
}
