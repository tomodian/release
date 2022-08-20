package flag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVersion(t *testing.T) {
	got := Version("/tmp")

	require.NotNil(t, got)
}
