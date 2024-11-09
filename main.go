package main

import (
	"log"
	"os"

	"github.com/marshallku/azutils/pkg/azure"
	"github.com/marshallku/azutils/pkg/config"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	if !azure.CheckCredential() {
		logger.Println("Azure credentials not found")
		os.Exit(1)
	}

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Println("Error loading config:", err)
		os.Exit(1)
	}

	logger.Println("Tags to keep:", cfg.TagsToKeep)
	logger.Println("Successfully logged in to Azure")
}
