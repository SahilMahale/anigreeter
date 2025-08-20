package db

import (
	"fmt"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type QuotesDb struct {
	gorm.Model
	Character string
	Anime     string
	QuoteText string
}

const DB_PATH = "./quotes/embeded.db"

var embeddedDB []byte

// SetEmbeddedDB set Embedded Db data
func SetEmbeddedDB(data []byte) {
	embeddedDB = data
}
func ExtractEmbeddedDB(dbPath string) error {

	if _, err := os.Stat(dbPath); err == nil {
		//skip extracting if it already exists
		return nil
	}

	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed while extracting DB:%v", err)
	}
	if err := os.WriteFile(dbPath, embeddedDB, 0644); err != nil {
		return fmt.Errorf("failed while writing and making the DB file: %v", err)
	}
	return nil
}
func EstablishConnectionEmbeddedDb() (*gorm.DB, error) {
	if err := ExtractEmbeddedDB(DB_PATH); err != nil {
		return nil, fmt.Errorf("failed in EstablishConnectionEmbeddedDb with: %w", err)
	}
	db, err := gorm.Open(sqlite.Open(DB_PATH))
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&QuotesDb{})
	return db, nil
}
