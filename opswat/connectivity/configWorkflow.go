package opswatClient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// GetConfigWorkflow - Returns workflow config
func (c *Client) GetConfigWorkflow() (*ConfigWorkflow, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/admin/config/rule", c.HostURL), nil)

	if err != nil {
		return nil, err
	}

	fmt.Println("request URL: " + fmt.Sprintf("%s/admin/config/rule", c.HostURL))

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := ConfigWorkflow{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateConfigWorkflow - Updates workflow config
func (c *Client) UpdateConfigWorkflow(config ConfigWorkflow) (*ConfigWorkflow, error) {

	preparedJson, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/config/rule", c.HostURL), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := ConfigWorkflow{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateConfigWorkflow - Creates workflow config
func (c *Client) CreateConfigWorkflow(config ConfigWorkflow) (*ConfigWorkflow, error) {

	preparedJson, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/admin/config/rule", c.HostURL), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := ConfigWorkflow{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteConfigWorkflow - Delete workflow config
func (c *Client) DeleteConfigWorkflow(configWorkflowID string) error {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/admin/config/rule/%s", c.HostURL, configWorkflowID), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return err
	}

	if string(body) != "{\"result\": \"Success\"}" {
		return errors.New(string(body))
	}

	return nil
}
