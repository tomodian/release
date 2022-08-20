package flag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewline(t *testing.T) {
	got := Newline("/tmp", true)

	require.NotNil(t, got)
}
