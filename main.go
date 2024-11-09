package main

import (
	"log"
	"os"

	"github.com/marshallku/azutils/cmd"
	"github.com/marshallku/azutils/pkg/azure"
	"github.com/spf13/cobra"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	if !azure.CheckCredential() {
		logger.Println("Azure credentials not found")
		os.Exit(1)
	}

	rootCmd := &cobra.Command{
		Use:   "azutils",
		Short: "Azure utilities for common tasks",
	}

	configCmd := cmd.NewConfigCommand(logger)
	rootCmd.AddCommand(configCmd)

	acrCmd := cmd.NewACRCommand(logger)
	rootCmd.AddCommand(acrCmd)

	if err := rootCmd.Execute(); err != nil {
		logger.Fatal(err)
	}
}
