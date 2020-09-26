package scaleway

import (
	"context"
	"fmt"
	"strings"

	Client "github.com/touchifyapp/cert-manager-webhook-scaleway/scaleway/client"

	"github.com/jetstack/cert-manager/pkg/acme/webhook/apis/acme/v1alpha1"
	"github.com/jetstack/cert-manager/pkg/issuer/acme/dns/util"
)

// GetBestDNSZone finds the best DNS Zone to use to add or update records
func (d *DNSClient) GetBestDNSZone(ch *v1alpha1.ChallengeRequest) (string, error) {
	orgID := Client.GoogleProtobufStringValue(d.organizationID)
	res, err := d.client.ListDNSZonesWithResponse(context.Background(), &Client.ListDNSZonesParams{OrganizationId: &orgID})
	if err != nil {
		return "", err
	}

	if res.StatusCode() < 200 || res.StatusCode() > 299 {
		return "", fmt.Errorf("invalid response status: %d\nbody: %s", res.StatusCode(), string(res.Body))
	}

	domain := util.UnFqdn(ch.ResolvedFQDN)
	result := ""

	for _, zone := range *res.JSON200.DnsZones {
		zoneDomain := createZoneDomain(*zone.Domain, *zone.Subdomain)
		if strings.HasSuffix(domain, zoneDomain) {
			if len(zoneDomain) > len(result) {
				result = zoneDomain
			}
		}
	}

	if result == "" {
		return "", fmt.Errorf("domain %s not found in DNS Zones", domain)
	}

	return result, nil
}

func createZoneDomain(domain string, subdomain string) string {
	if subdomain == "" {
		return domain
	}

	return subdomain + "." + domain
}
