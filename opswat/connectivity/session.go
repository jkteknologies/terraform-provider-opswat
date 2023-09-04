package opswatClient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetSession - Returns session config
func (c *Client) GetSession() (*Session, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/admin/config/session", c.HostURL), nil)

	if err != nil {
		return nil, err
	}

	fmt.Println("request URL: " + fmt.Sprintf("%s/admin/config/session", c.HostURL))

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := Session{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateSession - Updates session config
func (c *Client) UpdateSession(config Session) (*Session, error) {

	preparedJson, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/config/session", c.HostURL), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := Session{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateSession - Creates session config
func (c *Client) CreateSession(config Session) (*Session, error) {

	preparedJson, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/config/session", c.HostURL), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := Session{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
