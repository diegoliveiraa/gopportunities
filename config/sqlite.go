package config

import (
	"os"

	"github.com/diegoliveiraa/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	// Check if database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database file does not exist, creating a new one.")
		// Create the database file
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("Error creating database directory: %v", err)
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("Error creating database file: %v", err)
			return nil, err
		}
		file.Close()
		logger.Info("Database file created successfully.")
	}
	db, err := gorm.Open(
		sqlite.Open(dbPath),
		&gorm.Config{},
	)
	if err != nil {
		logger.Errorf("Error connecting to database: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Error migrating database: %v", err)
		return nil, err
	}
	return db, nil
}
