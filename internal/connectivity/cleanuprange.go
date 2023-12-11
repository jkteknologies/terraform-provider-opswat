package opswatClient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetScanHistory - Get processing history clean up time (clean up records older than).
func (c *Client) GetScanHistory() (*scanHistory, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/admin/config/scanhistory", c.HostURL), nil)

	if err != nil {
		return nil, err
	}

	fmt.Println("request URL: " + fmt.Sprintf("%s/admin/config/scanhistory", c.HostURL))

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := scanHistory{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateScanHistory - Update processing history clean up time (clean up records older than).
func (c *Client) UpdateScanHistory(cleanuprange int) (*scanHistory, error) {
	cleanuprangeJson := map[string]int{"cleanuprange": cleanuprange}
	preparedJson, err := json.Marshal(cleanuprangeJson)
	if err != nil {
		return nil, err
	}

	fmt.Println("----------- REQUEST -------------")
	fmt.Println(string(preparedJson), err)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/config/scanhistory", c.HostURL), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := scanHistory{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateScanHistory - Create processing history clean up time (clean up records older than).
func (c *Client) CreateScanHistory(cleanuprange int) (*scanHistory, error) {
	cleanuprangeJson := map[string]int{"cleanuprange": cleanuprange}
	preparedJson, err := json.Marshal(cleanuprangeJson)
	if err != nil {
		return nil, err
	}

	fmt.Println("----------- REQUEST -------------")
	fmt.Println(string(preparedJson), err)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/config/scanhistory", c.HostURL), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := scanHistory{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
