package quotegen

import (
	"fmt"
	"math/rand/v2"

	"github.com/SahilMahale/anigreeter/quotes"
)

type quoteGenService struct {
	anime     string
	character string
	seedEmbDb bool
}

func NewQuoteGenService(anime, character string, seedEmbDb bool) quoteGenService {
	return quoteGenService{
		anime:     anime,
		character: character,
		seedEmbDb: seedEmbDb,
	}
}

func (qg quoteGenService) GenerateQuote() error {
	if !qg.seedEmbDb {

		filteredQuotes, err := quotes.GetQuotes("anime")
		if err != nil {
			return err
		}

		if len(filteredQuotes) == 0 {
			return fmt.Errorf("no quotes found matching your criteria")
		}

		quote := filteredQuotes[rand.IntN(len(filteredQuotes))]
		fmt.Printf("\"%s\"\n  - %s (%s)\n", quote.QuoteText, quote.Character, quote.Anime)
	} else {
		fmt.Print("Seeding start....")

		err := quotes.SeedEmbeddedDb()
		if err != nil {
			return err
		}
		fmt.Printf("\nSeeding end\n")
	}
	return nil
}
