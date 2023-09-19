package opswatClient

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// GetWorkflows - Returns workflow configs
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

// GetWorkflow - Returns specific workflow configs
func (c *Client) GetWorkflow(workflowID int) (*Workflow, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/admin/config/rule/%d", c.HostURL, workflowID), nil)

	if err != nil {
		return nil, err
	}

	fmt.Println("request URL: " + fmt.Sprintf("%s/admin/config/rule", c.HostURL))

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := Workflow{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	//fmt.Println("UNMARSHAL RESULT")
	//fmt.Printf("Workflows : %+v", result)

	return &result, nil
}

// UpdateWorkflow - Updates workflow config
func (c *Client) UpdateWorkflow(workflowID int, workflow Workflow) (*Workflow, error) {

	preparedJson, err := json.Marshal(workflow)

	fmt.Println("----------- REQUEST -------------")
	fmt.Println(string(preparedJson), err)

	fmt.Println("----------- RULE ID -------------")
	fmt.Println(string(workflowID), err)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/config/rule/%d", c.HostURL, workflowID), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := Workflow{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateWorkflow - Creates workflow config
func (c *Client) CreateWorkflow(workflow Workflow) (*Workflow, error) {

	preparedJson, err := json.Marshal(workflow)

	fmt.Println("----------- REQUEST -------------")
	fmt.Println(string(preparedJson), err)

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

	result := Workflow{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteWorkflow - Delete workflow config
func (c *Client) DeleteWorkflow(workflowID string) error {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/admin/config/rule/%d", c.HostURL, workflowID), nil)
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
