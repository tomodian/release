package flag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDir(t *testing.T) {
	got := Dir("/tmp")

	require.NotNil(t, got)
}
