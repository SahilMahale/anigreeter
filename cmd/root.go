package cmd

import (
	"os"

	quotegen "github.com/SahilMahale/anigreeter/cmd/quoteGen"
	"github.com/spf13/cobra"
)

var (
	anime          string
	character      string
	seedEmbeddedDb bool
)

type quoteGenOps interface {
	GenerateQuote() error
}

var qgOPs quoteGenOps

var RootCmd = &cobra.Command{
	Use:   "anigreeter",
	Short: "A CLI tool for generating anime quotes",
	Long: `anigreeter is a command line tool that generates random anime quotes.
You can filter quotes by anime or character name.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		qgOPs = quotegen.NewQuoteGenService(anime, character, seedEmbeddedDb)
		return qgOPs.GenerateQuote()
	},
}

func init() {
	RootCmd.Flags().StringVarP(&anime, "anime", "a", "", "Filters quotes based on the anime they are from")
	RootCmd.Flags().StringVarP(&character, "character", "c", "", "Filters quotes based on the character that said it")
	if isDevMode() {
		RootCmd.Flags().BoolVarP(&seedEmbeddedDb, "seed", "s", false, "Seed the seed_embedded_db at once")
	}
}

func isDevMode() bool {
	return os.Getenv("ANIGREETER_MODE") == "dev"
}
