package main

import (
	"fmt"
	"github.com/khelaia/domingo/pkg/domingo"
	"github.com/khelaia/domingo/pkg/domingo/methods"
	"log"
)

func main() {
	client, err := domingo.NewClient()
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

	registerData, err := methods.RegisterDomain(client, "ditokhelaia11.com", "TXzRY#$EZ&o2;)B%[[4-npB8hNK0s,PP", "y", "1")
	if err != nil {
		log.Fatalf("Register Domain Failed: %v", err)
	}
	fmt.Println(registerData.Name, registerData.CreationDate, registerData.ExpirationDate)
	if err := methods.Logout(client); err != nil {
		log.Fatalf("Logout failed: %v", err)
	}
}
