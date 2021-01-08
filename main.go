package main

import (
	"os"

	"https://github.com/tensor-programming/golang-blockchain/tree/part_10/cli" //god fixed
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()
}
