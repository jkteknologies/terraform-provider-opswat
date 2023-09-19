package opswatClient

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
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

	/*tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}*/

	// WO for SR pc
	proxy, _ := url.Parse("http://gate-zrh.swissre.com:9443")

	c := Client{
		HTTPClient: &http.Client{Timeout: 1000 * time.Second, Transport: &http.Transport{Proxy: http.ProxyURL(proxy), TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}},
		//HTTPClient: &http.Client{Timeout: 100 * time.Second},
		//HTTPClient: &http.Client{Timeout: 100 * time.Second, Transport: tr},
		// Default OPSWAT URL
		HostURL: HostURL,
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

	//fmt.Println("apikey for req: " + c.Apikey)

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
