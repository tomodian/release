package flag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIgnoreEmpty(t *testing.T) {
	got := IgnoreEmpty("/tmp")

	require.NotNil(t, got)
}
