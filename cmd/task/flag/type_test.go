package flag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestType(t *testing.T) {
	got := Type("/tmp")

	require.NotNil(t, got)
}
