package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/khelaia/domingo/pkg/domingo"
	"github.com/khelaia/domingo/pkg/domingo/config"
	"github.com/khelaia/domingo/pkg/domingo/methods"
	"log"
	"os"
)

func main() {
	client, err := domingo.NewClient(loadConfig())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	if err := methods.Login(client); err != nil {
		log.Fatalf("Login failed: %v", err)
	}

	checkData, err := methods.CheckDomain(client, "ditokhelaia9.com")
	if err != nil {
		log.Fatalf("Check Domain failed: %v", err)
	}
	fmt.Println(checkData.Name, checkData.IsAvailable, checkData.Reason)

	/*
		registerData, err := methods.RegisterDomain(client, "ditokhelaia11.com", "TXzRY#$EZ&o2;)B%[[4-npB8hNK0s,PP", "y", "1")
		if err != nil {
			log.Fatalf("Register Domain Failed: %v", err)
		}
		fmt.Println(registerData.Name, registerData.CreationDate, registerData.ExpirationDate)
	*/

	if err := methods.Logout(client); err != nil {
		log.Fatalf("Logout failed: %v", err)
	}
}

func loadConfig() *config.EPPConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &config.EPPConfig{
		UserID:         os.Getenv("EPP_USERID"),
		Password:       os.Getenv("EPP_PASSWORD"),
		Hostname:       os.Getenv("EPP_HOSTNAME"),
		Port:           os.Getenv("EPP_PORT"),
		ClientCertFile: os.Getenv("EPP_CLIENT_CERT_FILE"),
		ClientKeyFile:  os.Getenv("EPP_CLIENT_KEY_FILE"),
	}
}
