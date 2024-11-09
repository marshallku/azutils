package resources

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"

	"github.com/marshallku/azutils/pkg/config"
)

func GetRepositories(registry *string) ([]string, error) {
	if registry == nil {
		config, err := config.NewConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to get config: %v", err)
		}
		registry = &config.Registry
	}

	cmd := exec.Command("az", "acr", "repository", "list", "--name", *registry)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to list repositories: %v, output: %s", err, string(output))
	}

	var repositories []string
	if err := json.Unmarshal(output, &repositories); err != nil {
		return nil, fmt.Errorf("failed to parse repositories JSON: %v", err)
	}

	return repositories, nil
}

func GetTags(imageName string, registry *string) ([]string, error) {
	if registry == nil {
		config, err := config.NewConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to get config: %v", err)
		}
		registry = &config.Registry
	}

	cmd := exec.Command("az", "acr", "repository", "show-tags", "--name", *registry, "--repository", imageName, "--orderby", "time_desc")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to get tags for repository %s: %v, output: %s", imageName, err, string(output))
	}

	var tags []string
	if err := json.Unmarshal(output, &tags); err != nil {
		return nil, fmt.Errorf("failed to parse tags JSON: %v", err)
	}

	return tags, nil
}

func RemoveTags(imageName string, tagsToKeep *int, registry *string) error {
	config, err := config.NewConfig()
	if err != nil {
		return fmt.Errorf("failed to get config: %v", err)
	}

	if registry == nil {
		registry = &config.Registry
	}

	if tagsToKeep == nil {
		tagsToKeep = &config.TagsToKeep
	}

	tags, err := GetTags(imageName, registry)
	if err != nil {
		return err
	}

	if len(tags) <= *tagsToKeep {
		fmt.Printf("No tags to delete for %s, as there are %v or fewer\n", imageName, *tagsToKeep)
		return nil
	}

	for _, tag := range tags[*tagsToKeep:] {
		image := fmt.Sprintf("%s:%s", imageName, tag)
		cmd := exec.Command("az", "acr", "repository", "delete", "--name", *registry, "--image", image, "--yes")
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("Failed to delete tag %s:%v\noutput: %s", image, err, string(output))
		} else {
			fmt.Printf("Deleted tag: %s\n", tag)
		}
	}

	return nil
}
