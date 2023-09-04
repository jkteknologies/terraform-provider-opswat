package opswatClient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// GetWorkflows - Returns workflows config
func (c *Client) GetWorkflows() (*Workflows, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/admin/config/rule", c.HostURL), nil)

	if err != nil {
		return nil, err
	}

	fmt.Println("request URL: " + fmt.Sprintf("%s/admin/config/rule", c.HostURL))

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := Workflows{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateWorkflow - Updates workflow config
func (c *Client) UpdateWorkflow(config Workflows) (*Workflows, error) {

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

	result := Workflows{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateWorkflow - Creates workflow config
func (c *Client) CreateWorkflow(config Workflows) (*Workflows, error) {

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

	result := Workflows{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteWorkflow - Delete workflow config
func (c *Client) DeleteWorkflow(configWorkflowID string) error {

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
