package main

import (
	"os"

	"github.com/Redstone-2-0/core-in-go/cli" //god fixed
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()
}
