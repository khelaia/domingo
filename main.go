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
		//Register Domain
		registerData, err := methods.RegisterDomain(client, "ditokhelaia12.com", "TXzRY#$EZ&o2;)B%[[4-npB8hNK0s,PP", "y", "1")
		if err != nil {
			log.Fatalf("Register Domain Failed: %v", err)
		}
		fmt.Println(registerData.Name, registerData.CreationDate, registerData.ExpirationDate)
	*/
	/*
		//Create Nameservers
		createHost, err := methods.CreateHost(client, "ns5.ditokhelaia11.com", "104.21.29.230")
		if err != nil {
			log.Fatalf("Failed to create host %s", err)
		}
		fmt.Println(createHost.Message, createHost.HostName, createHost.CreationDate)
	*/

	/*
		nameservers := []string{"ns1.ditokhelaia11.com", "ns2.ditokhelaia11.com", "ns3.ditokhelaia11.com"}
		attachData, err := methods.AttachNameservers(client, "ditokhelaia11.com", nameservers)
		if err != nil {
			log.Fatalf("Failed to attach host to domain %s", err)
		}
		msg := *attachData

		fmt.Println(msg)
	*/

	/*
		statuses := []constants.ClientStatus{constants.StatusClientUpdateProhibited, constants.StatusClientRenewProhibited}
		addStatusData, err := methods.AddStatuses(client, "ditokhelaia11.com", statuses)
		if err != nil {
			log.Fatalf("Failed to add statuses to domain %s", err)
		}
		msg := *addStatusData
		fmt.Println(msg)
	*/

	infoData, err := methods.DomainInfo(client, "ditokhelaia11.com")
	if err != nil {
		log.Fatalf("Failed to fetch domain info %s", err)
	}

	fmt.Printf("%+v\n", infoData)

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
