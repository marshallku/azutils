package main

import (
	"log"
	"os"

	"github.com/marshallku/azutils/pkg/azure"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	client := azure.NewAzureClient(nil)

	if !client.CheckCredential() {
		client.LoginWithEnvironmentalVariables()
	}

	logger.Println("Successfully logged in to Azure")
}
