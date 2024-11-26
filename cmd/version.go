package cmd

import (
	"fmt"

	"github.com/marshallku/azutils/pkg/version"
	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version",
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println()
			fmt.Printf(" ░▒▓██████▓▒░░▒▓████████▓▒░▒▓█▓▒░░▒▓█▓▒░▒▓████████▓▒░▒▓█▓▒░▒▓█▓▒░       ░▒▓███████▓▒░ \n")
			fmt.Printf("░▒▓█▓▒░░▒▓█▓▒░      ░▒▓█▓▒░▒▓█▓▒░░▒▓█▓▒░  ░▒▓█▓▒░   ░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░        \n")
			fmt.Printf("░▒▓█▓▒░░▒▓█▓▒░    ░▒▓██▓▒░░▒▓█▓▒░░▒▓█▓▒░  ░▒▓█▓▒░   ░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░        \n")
			fmt.Printf("░▒▓████████▓▒░  ░▒▓██▓▒░  ░▒▓█▓▒░░▒▓█▓▒░  ░▒▓█▓▒░   ░▒▓█▓▒░▒▓█▓▒░       ░▒▓██████▓▒░  \n")
			fmt.Printf("░▒▓█▓▒░░▒▓█▓▒░░▒▓██▓▒░    ░▒▓█▓▒░░▒▓█▓▒░  ░▒▓█▓▒░   ░▒▓█▓▒░▒▓█▓▒░             ░▒▓█▓▒░ \n")
			fmt.Printf("░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░  ░▒▓█▓▒░   ░▒▓█▓▒░▒▓█▓▒░             ░▒▓█▓▒░ \n")
			fmt.Printf("░▒▓█▓▒░░▒▓█▓▒░▒▓████████▓▒░░▒▓██████▓▒░   ░▒▓█▓▒░   ░▒▓█▓▒░▒▓████████▓▒░▒▓███████▓▒░  \n")
			fmt.Println()

			fmt.Printf("Version: %s\n", version.Version)
			fmt.Printf("Commit : %s\n", version.Commit)
			fmt.Printf("Date   : %s\n", version.Date)
		},
	}

	return cmd
}
