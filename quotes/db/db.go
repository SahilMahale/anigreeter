package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type QuotesDb struct {
	gorm.Model
	Character string
	Anime     string
	QuoteText string
}

func EstablishConnectionEmbeddedDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./quotes/embeded.db"))
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&QuotesDb{})
	return db, nil
}
