package opswatClient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// GetDir - Returns global sync scan timeout
func (c *Client) GetDir() (*userDirectory, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/admin/userdirectory", c.HostURL), nil)

	if err != nil {
		return nil, err
	}

	fmt.Println("request URL: " + fmt.Sprintf("%s/admin/userdirectory", c.HostURL))

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := userDirectory{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateDir - Update global sync scan timeout
func (c *Client) UpdateDir(timeout int) (*userDirectory, error) {
	timeoutJson := map[string]int{"timeout": timeout}
	preparedJson, err := json.Marshal(timeoutJson)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/userdirectory", c.HostURL), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := userDirectory{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateDir - Creates global sync scan timeout
func (c *Client) CreateDir(timeout int) (*userDirectory, error) {
	timeoutJson := map[string]int{"timeout": timeout}
	preparedJson, err := json.Marshal(timeoutJson)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/userdirectory", c.HostURL), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := userDirectory{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteUserDirectory - Delete userdirectory
func (c *Client) DeleteUserDirectory(dirID int) error {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/admin/userdirectory/%d", c.HostURL, dirID), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return err
	}

	if string(body) != `{"result":"Success"}` {
		return errors.New(string(body))
	}

	return nil
}
