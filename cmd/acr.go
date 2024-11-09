package cmd

import (
	"fmt"
	"log"

	"github.com/marshallku/azutils/pkg/resources"
	"github.com/spf13/cobra"
)

func NewACRCommand(logger *log.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "acr",
		Short: "Azure Container Registry operations",
		Long:  `Commands to manage Azure Container Registry repositories and tags`,
	}

	listRepoCmd := &cobra.Command{
		Use:   "list-repos",
		Short: "List all repositories in ACR",
		RunE: func(cmd *cobra.Command, args []string) error {
			registry, _ := cmd.Flags().GetString("registry")
			var registryPtr *string
			if registry != "" {
				registryPtr = &registry
			}

			repositories, err := resources.GetRepositories(registryPtr)
			if err != nil {
				return err
			}

			for _, repo := range repositories {
				fmt.Println(repo)
			}
			return nil
		},
	}

	listTagsCmd := &cobra.Command{
		Use:   "list-tags [image-name]",
		Short: "List all tags for a repository",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			imageName := args[0]
			registry, _ := cmd.Flags().GetString("registry")
			var registryPtr *string
			if registry != "" {
				registryPtr = &registry
			}

			tags, err := resources.GetTags(imageName, registryPtr)
			if err != nil {
				return err
			}

			for _, tag := range tags {
				fmt.Println(tag)
			}
			return nil
		},
	}

	removeTagsCmd := &cobra.Command{
		Use:   "remove-tags [image-name]",
		Short: "Remove old tags from a repository",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			imageName := args[0]
			registry, _ := cmd.Flags().GetString("registry")
			tagsToKeep, _ := cmd.Flags().GetInt("keep")

			var registryPtr *string
			if registry != "" {
				registryPtr = &registry
			}

			var tagsToKeepPtr *int
			if cmd.Flags().Changed("keep") {
				tagsToKeepPtr = &tagsToKeep
			}

			return resources.RemoveTags(imageName, tagsToKeepPtr, registryPtr)
		},
	}

	pruneTagsCmd := &cobra.Command{
		Use:   "prune-tags",
		Short: "Remove all tags from a repository",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			imageName, _ := cmd.Flags().GetString("image-name")
			registry, _ := cmd.Flags().GetString("registry")
			tagsToKeep, _ := cmd.Flags().GetInt("keep")

			var registryPtr *string
			if registry != "" {
				registryPtr = &registry
			}

			var tagsToKeepPtr *int
			if cmd.Flags().Changed("keep") {
				tagsToKeepPtr = &tagsToKeep
			}

			if imageName != "" {
				return resources.RemoveTags(imageName, tagsToKeepPtr, registryPtr)
			}

			repositories, err := resources.GetRepositories(registryPtr)
			if err != nil {
				return err
			}

			for _, repo := range repositories {
				fmt.Printf("Pruning tags for repository %s\n", repo)
				err := resources.RemoveTags(repo, tagsToKeepPtr, registryPtr)
				if err != nil {
					return err
				}
			}

			return fmt.Errorf("repository %s not found", imageName)
		},
	}

	listRepoCmd.Flags().String("registry", "", "ACR registry name (optional)")
	listTagsCmd.Flags().String("registry", "", "ACR registry name (optional)")
	removeTagsCmd.Flags().String("registry", "", "ACR registry name (optional)")
	removeTagsCmd.Flags().Int("keep", 0, "Number of tags to keep (optional)")
	pruneTagsCmd.Flags().String("image-name", "", "Image name to prune tags from (optional)")
	pruneTagsCmd.Flags().String("registry", "", "ACR registry name (optional)")
	pruneTagsCmd.Flags().Int("keep", 0, "Number of tags to keep (optional)")

	cmd.AddCommand(listRepoCmd)
	cmd.AddCommand(listTagsCmd)
	cmd.AddCommand(removeTagsCmd)
	cmd.AddCommand(pruneTagsCmd)

	return cmd
}
