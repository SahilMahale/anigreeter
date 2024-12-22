package quotegen

import (
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/SahilMahale/anigreeter/quotes"
)

type quoteGenService struct {
	anime     string
	character string
}

func NewQuoteGenService(anime, character string) quoteGenService {
	return quoteGenService{
		anime:     anime,
		character: character,
	}
}

func (qg quoteGenService) GenerateQuote() error {
	allQuotes := quotes.GetQuotes()
	filteredQuotes := qg.filterQuotes(allQuotes)

	if len(filteredQuotes) == 0 {
		return fmt.Errorf("no quotes found matching your criteria")
	}

	quote := filteredQuotes[rand.IntN(len(filteredQuotes))]
	fmt.Printf("\"%s\"\n  - %s (%s)\n", quote.Text, quote.Character, quote.Anime)
	return nil
}

func (qg quoteGenService) filterQuotes(qts []quotes.Quote) []quotes.Quote {
	filtered := []quotes.Quote{}

	for _, q := range qts {
		if q.Anime != "" && !strings.Contains(strings.ToLower(q.Anime), strings.ToLower(qg.anime)) {
			continue
		}
		if q.Character != "" && !strings.Contains(strings.ToLower(q.Character), strings.ToLower(qg.character)) {
			continue
		}
		filtered = append(filtered, q)
	}

	return filtered
}
