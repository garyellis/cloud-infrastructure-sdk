package main

import (
	"os"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/cli"
)

func main() {
	if err := cli.NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
