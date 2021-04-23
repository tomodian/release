package parser_test

import (
	"testing"

	"release/parser"

	"github.com/stretchr/testify/assert"
)

func TestVersionTypeString(t *testing.T) {
	assert.Equal(t, string(parser.MajorVersion), parser.MajorVersion.String())
}
