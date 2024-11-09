package azure

import (
	"encoding/json"
	"os"
	"os/exec"
)

type Client interface {
	CheckCredential() bool
	ListRepositories(registry string) ([]string, error)
	ListTags(registry, repository string) ([]string, error)
	DeleteTag(registry, repository, tag string) error
}

type AzureCredentials struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	TenantID     string `json:"tenantId"`
}

func CheckCredential() bool {
	cmd := exec.Command("az", "account", "show")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}

	var account struct {
		User struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"user"`
	}

	if err := json.Unmarshal(output, &account); err != nil {
		return false
	}

	return account.User.Name != ""
}

func LoginWithEnvironmentalVariables() error {
	credentialJson := os.Getenv("AZURE_CREDENTIALS")

	if credentialJson != "" {
		var cfg AzureCredentials
		if err := json.Unmarshal([]byte(credentialJson), &cfg); err != nil {
			return err
		}

		return Login(cfg)
	}

	clientId := os.Getenv("AZURE_CLIENT_ID")
	clientSecret := os.Getenv("AZURE_CLIENT_SECRET")
	tenantId := os.Getenv("AZURE_TENANT_ID")

	if clientId == "" || clientSecret == "" || tenantId == "" {
		return nil
	}

	return Login(AzureCredentials{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TenantID:     tenantId,
	})
}

func Login(creds AzureCredentials) error {
	cmd := exec.Command("az", "login", "--service-principal",
		"-u", creds.ClientID,
		"-p", creds.ClientSecret,
		"--tenant", creds.TenantID)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}
