package scaleway

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func getSubDomain(zone string, domain string) string {
	if idx := strings.Index(domain, "."+zone); idx != -1 {
		return domain[:idx]
	}

	return ""
}

func validateResponse(err error, res *http.Response) error {
	if err != nil {
		return err
	}

	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}

	body, err := readResponse(res)
	if err != nil {
		return err
	}

	return fmt.Errorf("invalid response status: %d\nbody: %s", res.StatusCode, body)
}

func readResponse(res *http.Response) (string, error) {
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
