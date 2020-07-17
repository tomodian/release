package utils_test

import (
	"testing"

	"release/utils"

	"github.com/stretchr/testify/assert"
)

func TestPretty(t *testing.T) {
	type pattern struct {
		expected string
		input    string
	}

	pats := []pattern{
		{
			expected: utils.EmptyLine,
			input:    utils.EmptyLine,
		},
		{
			expected: "Title",
			input:    "### Title",
		},
		{
			expected: "hello",
			input:    "hello",
		},
	}

	for _, p := range pats {
		got := utils.Pretty(p.input)

		assert.NotEmpty(t, got)
		assert.Contains(t, got, p.expected)
	}
}
