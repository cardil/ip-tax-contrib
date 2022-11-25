package main_test

import (
	"bytes"
	"math"
	"testing"

	mainapp "github.com/cardil/ip-tax-contributions"
	"github.com/cardil/ip-tax-contributions/cmd"
	"github.com/stretchr/testify/assert"
	"github.com/wavesoftware/go-commandline"
)

func TestMainFunc(t *testing.T) {
	retcode := math.MinInt64
	defer func() {
		cmd.Options = nil
	}()
	var buf bytes.Buffer
	cmd.Options = []commandline.Option{
		commandline.WithExit(func(code int) {
			retcode = code
		}),
		commandline.WithOutput(&buf),
		commandline.WithArgs("--help"),
	}

	mainapp.Main()

	out := buf.String()
	assert.Contains(t, out, "Generate a list of contributions for 🇵🇱 IP Tax form")
	assert.Equal(t, retcode, math.MinInt64)
}
