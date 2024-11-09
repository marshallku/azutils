package cmd

import (
	"fmt"
	"log"

	"github.com/marshallku/azutils/pkg/config"
	"github.com/spf13/cobra"
)

func NewConfigCommand(logger *log.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Configuration operations",
		Long:  `Commands to manage configuration settings`,
	}

	updateConfigCommand := &cobra.Command{
		Use:   "update [key] [value]",
		Short: "Update a configuration setting",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			key := args[0]
			value := args[1]
			if key == "" || value == "" {
				return fmt.Errorf("key and value must be provided")
			}

			// Capitalize first letter of key
			key = string(key[0]-32) + key[1:]

			if err := config.UpdateConfig(key, value); err != nil {
				return err
			}

			logger.Printf("Updated configuration %s to %s\n", key, value)
			return nil
		},
	}

	cmd.AddCommand(updateConfigCommand)

	return cmd
}
