package opswatClient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetGlobalSync - Returns global sync scan timeout
func (c *Client) GetConfigSession() (*configSession, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/admin/config/file/sync", c.HostURL), nil)

	if err != nil {
		return nil, err
	}

	fmt.Println("request URL: " + fmt.Sprintf("%s/admin/config/file/sync", c.HostURL))

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := configSession{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateGlobalSync - Returns global sync scan timeout
func (c *Client) UpdateConfigSession(timeout int) (*configSession, error) {
	timeoutJson := map[string]int{"timeout": timeout}
	preparedJson, err := json.Marshal(timeoutJson)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/config/file/sync", c.HostURL), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := configSession{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateGlobalSync - Returns global sync scan timeout
func (c *Client) CreateConfigSession(timeout int) (*configSession, error) {
	timeoutJson := map[string]int{"timeout": timeout}
	preparedJson, err := json.Marshal(timeoutJson)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/config/file/sync", c.HostURL), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := configSession{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
