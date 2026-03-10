package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomodian/release/cmd"
)

func TestRun(t *testing.T) {
	assert.NotPanics(t, func() {
		_ = cmd.Run(os.Args[0:1])
	})
}
