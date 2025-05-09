package config

import (
	"gorm.io/gorm"
)

var (
	// DB is the database connection
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return nil
}

func GetLogger(p string) *Logger {

	logger = NewLogger(p)
	return logger
}
