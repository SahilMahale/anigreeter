package main

import (
	"fmt"
	"os"

	"github.com/SahilMahale/anigreeter/src/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
