package opswatClient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// GetWorkflows - Returns workflows config
func (c *Client) GetWorkflows() ([]Workflow, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/admin/config/rule", c.HostURL), nil)

	if err != nil {
		return nil, err
	}

	fmt.Println("request URL: " + fmt.Sprintf("%s/admin/config/rule", c.HostURL))

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := []Workflow{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	//fmt.Println("UNMARSHAL RESULT")
	//fmt.Printf("Workflows : %+v", result)

	return result, nil
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
