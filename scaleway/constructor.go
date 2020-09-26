package scaleway

import (
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	Client "github.com/touchifyapp/cert-manager-webhook-scaleway/scaleway/client"
)

// DNSClient implements Scaleway utilities functions to request and create DNS records.
type DNSClient struct {
	organizationID string
	client         *Client.ClientWithResponses
}

// NewClient creates a new OpenAPI Client to connect Scaleway DNS API
func NewClient(organizationID string, secretKey string) (*DNSClient, error) {
	// Example ApiKey provider
	// See: https://swagger.io/docs/specification/authentication/api-keys/
	apiKeyProvider, apiKeyProviderErr := securityprovider.NewSecurityProviderApiKey("header", "X-Auth-Token", secretKey)
	if apiKeyProviderErr != nil {
		return nil, apiKeyProviderErr
	}

	client, clientErr := Client.NewClientWithResponses(
		"https://api.scaleway.com",
		Client.WithBaseURL("https://api.scaleway.com"),
		Client.WithRequestEditorFn(apiKeyProvider.Intercept),
	)

	if clientErr != nil {
		return nil, clientErr
	}

	dnsClient := DNSClient{}
	dnsClient.organizationID = organizationID
	dnsClient.client = client

	return &dnsClient, nil
}
