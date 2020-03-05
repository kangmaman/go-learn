package main

import (
	"errors"

	"github.com/joho/godotenv"
)

//Config .... struct config handle konfigurasi
type Config struct {
}

//LoadConfig ... method config untuk handle konfigurasi
func (cfg *Config) LoadConfig() error {
	if err := godotenv.Load(); err != nil {
		return errors.New("No .env file found")
	}
	return nil
}
