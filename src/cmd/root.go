package cmd

import (
	quotegen "github.com/SahilMahale/anigreeter/src/cmd/quoteGen"
	"github.com/spf13/cobra"
)

var (
	anime     string
	character string
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
		return qgOPs.GenerateQuote()
	},
}

func init() {
	RootCmd.Flags().StringVarP(&anime, "anime", "a", "", "Filters quotes based on the anime they are from")
	RootCmd.Flags().StringVarP(&character, "character", "c", "", "Filters quotes based on the character that said it")
	qgOPs = quotegen.NewQuoteGenService(anime, character)
}
