package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/SahilMahale/anigreeter/cmd"
)

//go:embed quotes/embeded.db
var embeddedDB []byte

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
