package main

import (
	"github.com/cardil/ip-tax-contributions/cmd"
	"github.com/wavesoftware/go-commandline"
)

func main() {
	commandline.New(new(cmd.App)).ExecuteOrDie(cmd.Options...)
}

// Main is used for testing purposes.
func Main() { //nolint:deadcode
	main()
}
