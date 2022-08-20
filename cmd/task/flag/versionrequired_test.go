package flag

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVersionRequired(t *testing.T) {
	got := VersionRequired("/tmp")

	require.NotNil(t, got)
}
