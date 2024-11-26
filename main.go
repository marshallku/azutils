package main

import (
	"log"
	"os"

	"github.com/marshallku/azutils/cmd"
	"github.com/marshallku/azutils/pkg/azure"
	v "github.com/marshallku/azutils/pkg/version"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	v.Version = version
	v.Commit = commit
	v.Date = date

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

	versionCmd := cmd.NewVersionCommand()
	rootCmd.AddCommand(versionCmd)

	acrCmd := cmd.NewACRCommand(logger)
	rootCmd.AddCommand(acrCmd)

	if err := rootCmd.Execute(); err != nil {
		logger.Fatal(err)
	}
}
