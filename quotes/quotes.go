package quotes

import (
	"fmt"

	"github.com/SahilMahale/anigreeter/quotes/db"
)

var (
	// INFO: List supported filters
	ANIME      = "anime"
	CHARACTER  = "character"
	FULLRANDOM = "fullRandom"
)

// SeedEmbeddedDb Helps devs quickly seed the embedded Db, only work in dev mode,
// set ANIGREETER_MODE=dev
func SeedEmbeddedDb() error {
	db, err := db.EstablishConnectionEmbeddedDb()
	if err != nil {
		return err
	}
	for _, qt := range GetQuotesToEmbed() {
		res := db.Create(&qt)
		if res.Error != nil {
			return res.Error
		}
		fmt.Print(".")
	}
	return nil
}

// GetQuotes takes in a filter and returns filtered quotes array,
// supported filters [anime,character or fullRandom]
func GetQuotes(filter string) ([]db.QuotesDb, error) {
	quotes := []db.QuotesDb{}
	switch filter {
	case ANIME:
		break
	case CHARACTER:
		break
	case FULLRANDOM:
		break
	default:
		return nil, fmt.Errorf("filter not supported")
	}
	return quotes, nil
}

// GetQuotesToEmbed just returns a formated array of quotes to feed into
// the seeding function
func GetQuotesToEmbed() []db.QuotesDb {
	quotes := []db.QuotesDb{
		// Insert entries to be bulk embedded
	}

	return quotes
}
