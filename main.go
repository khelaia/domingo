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

	data, err := methods.CheckDomain(client, "ditokhelaia9.com")
	if err != nil {
		log.Fatalf("Check Domain failed: %v", err)
	}
	fmt.Println(data.Name, data.IsAvailable, data.Reason)

	if err := methods.Logout(client); err != nil {
		log.Fatalf("Logout failed: %v", err)
	}
}
