package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type EPPConfig struct {
	UserID         string
	Password       string
	Hostname       string
	Port           string
	ClientCertFile string
	ClientKeyFile  string
}

func LoadConfig() *EPPConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &EPPConfig{
		UserID:         os.Getenv("EPP_USERID"),
		Password:       os.Getenv("EPP_PASSWORD"),
		Hostname:       os.Getenv("EPP_HOSTNAME"),
		Port:           os.Getenv("EPP_PORT"),
		ClientCertFile: os.Getenv("EPP_CLIENT_CERT_FILE"),
		ClientKeyFile:  os.Getenv("EPP_CLIENT_KEY_FILE"),
	}
}
