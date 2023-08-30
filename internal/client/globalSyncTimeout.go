package opswatClient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetGlobalSync - Returns global sync scan timeout
func (c *Client) GetGlobalSync() (*globalSyncTimeout, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/admin/config/file/sync", c.HostURL), nil)

	if err != nil {
		return nil, err
	}

	fmt.Println("request URL: " + fmt.Sprintf("%s/admin/config/file/sync", c.HostURL))

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := globalSyncTimeout{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateGlobalSync - Returns global sync scan timeout
func (c *Client) UpdateGlobalSync() (*globalSyncTimeout, error) {
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/config/file/sync", c.HostURL), nil)

	if err != nil {
		return nil, err
	}

	fmt.Println("request URL: " + fmt.Sprintf("%s/admin/config/file/sync", c.HostURL))

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := globalSyncTimeout{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
