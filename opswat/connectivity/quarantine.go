package opswatClient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetQuarantine - Returns session config
func (c *Client) GetQuarantine() (*Quarantine, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/admin/config/quarantine", c.HostURL), nil)

	if err != nil {
		return nil, err
	}

	fmt.Println("request URL: " + fmt.Sprintf("%s/admin/config/quarantine", c.HostURL))

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := Quarantine{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateQuarantine - Updates session config
func (c *Client) UpdateQuarantine(config Quarantine) (*Quarantine, error) {

	preparedJson, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/config/quarantine", c.HostURL), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := Quarantine{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateQuarantine - Creates session config
func (c *Client) CreateQuarantine(config Quarantine) (*Quarantine, error) {

	preparedJson, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/config/quarantine", c.HostURL), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := Quarantine{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
