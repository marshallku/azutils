package azure

import (
	"encoding/json"
	"os"
	"os/exec"

	"github.com/marshallku/azutils/pkg/config"
)

type Client interface {
	CheckCredential() bool
	Login(credentials config.AzureCredentials) error
	ListRepositories(registry string) ([]string, error)
	ListTags(registry, repository string) ([]string, error)
	DeleteTag(registry, repository, tag string) error
}

type AzureCredentials struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	TenantID     string `json:"tenantId"`
}

type AzureClient struct {
	executor CommandExecutor
}

type CommandExecutor interface {
	Execute(name string, args ...string) ([]byte, error)
}

type DefaultExecutor struct{}

func (e *DefaultExecutor) Execute(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	return cmd.CombinedOutput()
}

func NewAzureClient(executor CommandExecutor) *AzureClient {
	if executor == nil {
		executor = &DefaultExecutor{}
	}
	return &AzureClient{executor: executor}
}

func (c *AzureClient) CheckCredential() bool {
	output, err := c.executor.Execute("az", "account", "show")
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

func (c *AzureClient) LoginWithEnvironmentalVariables() error {
	credentialJson := os.Getenv("AZURE_CREDENTIALS")

	if credentialJson != "" {
		var cfg AzureCredentials
		if err := json.Unmarshal([]byte(credentialJson), &cfg); err != nil {
			return err
		}

		return c.Login(cfg)
	}

	clientId := os.Getenv("AZURE_CLIENT_ID")
	clientSecret := os.Getenv("AZURE_CLIENT_SECRET")
	tenantId := os.Getenv("AZURE_TENANT_ID")

	if clientId == "" || clientSecret == "" || tenantId == "" {
		return nil
	}

	return c.Login(AzureCredentials{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TenantID:     tenantId,
	})
}

func (c *AzureClient) Login(creds AzureCredentials) error {
	_, err := c.executor.Execute("az", "login", "--service-principal",
		"-u", creds.ClientID,
		"-p", creds.ClientSecret,
		"--tenant", creds.TenantID)
	return err
}
