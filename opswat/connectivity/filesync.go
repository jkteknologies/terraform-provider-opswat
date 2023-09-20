package opswatClient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

// UpdateGlobalSync - Update global sync scan timeout
func (c *Client) UpdateGlobalSync(timeout int) (*globalSyncTimeout, error) {
	timeoutJson := map[string]int{"timeout": timeout}
	preparedJson, err := json.Marshal(timeoutJson)
	if err != nil {
		return nil, err
	}

	fmt.Println("----------- REQUEST -------------")
	fmt.Println(string(preparedJson), err)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/config/file/sync", c.HostURL), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

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

// CreateGlobalSync - Creates global sync scan timeout
func (c *Client) CreateGlobalSync(timeout int) (*globalSyncTimeout, error) {
	timeoutJson := map[string]int{"timeout": timeout}
	preparedJson, err := json.Marshal(timeoutJson)
	if err != nil {
		return nil, err
	}

	fmt.Println("----------- REQUEST -------------")
	fmt.Println(string(preparedJson), err)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/config/file/sync", c.HostURL), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

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
