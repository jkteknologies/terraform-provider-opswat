package opswatClient

import (
	"fmt"
	"io"
	"net/http"
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

// NewClient -
func NewClient(host, apikey *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default OPSWAT URL
		HostURL: "https://" + HostURL,
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
