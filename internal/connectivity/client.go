package opswatClient

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

// HostURL - Default OPSWAT API URL
const HostURL string = "localhost"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Apikey     string
}

// NewClient - test
func NewClient(host, apikey *string) (*Client, error) {

	// WO for corporate proxy
	proxy, _ := url.Parse("https://localhost:8080")
	tr := &http.Transport{}
	val, present := os.LookupEnv("HTTPS_PROXY")
	if present {
		proxy, _ = url.Parse(val)
		tr = &http.Transport{Proxy: http.ProxyURL(proxy), TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	} else {
		tr = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	}

	c := Client{
		HTTPClient: &http.Client{Timeout: 1000 * time.Second, Transport: tr},
		HostURL:    HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	if apikey != nil {
		c.Apikey = *apikey
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {

	req.Header.Set("apikey", c.Apikey)

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
