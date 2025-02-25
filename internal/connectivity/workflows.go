package opswatClient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// GetWorkflows - Returns workflow configs
func (c *Client) GetWorkflows(ctx context.Context) ([]Workflow, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/admin/config/rule", c.HostURL), nil)
	
	if err != nil {
		return nil, err
	}
	
	tflog.Debug(ctx, "request URL: " + fmt.Sprintf("%s/admin/config/rule", c.HostURL))
	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	// Opswat uses '#' symbol as All roles marker, need to convert it to 0
	bodyNormalized := NormalizeWorkflows(ctx, body)

	result := []Workflow{}

	err = json.Unmarshal(bodyNormalized, &result)
	if err != nil {
		return nil, err
	}

	//fmt.Println("UNMARSHAL RESULT")
	//fmt.Printf("Workflows : %+v", result)

	return result, nil
}

// GetWorkflow - Returns specific workflow configs
func (c *Client) GetWorkflow(ctx context.Context, workflowID int) (*Workflow, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/admin/config/rule/%d", c.HostURL, workflowID), nil)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "request URL: " + fmt.Sprintf("%s/admin/config/rule/%d", c.HostURL, workflowID))
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

// UpdateWorkflow - Updates workflow config
func (c *Client) UpdateWorkflow(ctx context.Context, workflowID int, workflow Workflow) (*Workflow, error) {
	preparedJson, err := json.Marshal(workflow)

	if err != nil {
		return nil, err
	}
	
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/config/rule/%d", c.HostURL, workflowID), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}
	
	tflog.Debug(ctx, "request URL: " + fmt.Sprintf("%s/admin/config/rule/%d, request body: %s", c.HostURL, workflowID, string(preparedJson)))
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
func (c *Client) CreateWorkflow(ctx context.Context, workflow Workflow) (*Workflow, error) {
	preparedJson, err := json.Marshal(workflow)

	if err != nil {
		return nil, err
	}
	
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/admin/config/rule", c.HostURL), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}
	
	tflog.Debug(ctx, "request URL: " + fmt.Sprintf("%s/admin/config/rule, request body: %s", c.HostURL, string(preparedJson)))
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
func (c *Client) DeleteWorkflow(ctx context.Context, workflowID int) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/admin/config/rule/%d", c.HostURL, workflowID), nil)
	if err != nil {
		return err
	}
	
	tflog.Debug(ctx, "request URL: " + fmt.Sprintf("%s/admin/config/rule/%d", c.HostURL, workflowID))
	body, err := c.doRequest(req)

	if err != nil {
		return err
	}

	if string(body) != `{"result":"Success"}` {
		return errors.New(string(body))
	}

	return nil
}

func NormalizeWorkflows(ctx context.Context, jsonData []byte) []byte {
	var rawData []map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &rawData); err != nil {
		tflog.Info(ctx, "Error processing JSON response from the server")
	}

	// Normalize the "role" field
	for _, item := range rawData {
		tflog.Warn(ctx, "looping over raw data")
		if resultsAllowed, exists := item["result_allowed"]; exists {
			tflog.Warn(ctx, "in result_allowed")
			results, ok := resultsAllowed.([]interface{})
			if !ok {
				tflog.Error(ctx, "Unexpected type of server response")
			}
			
			// Iterate through each result in result_allowed
			for _, result := range results {
				resultMap, ok := result.(map[string]interface{})
				if !ok {
					tflog.Error(ctx, "Unexpected type of server response")
				}
				
				// Normalize the "role" field
				if role, exists := resultMap["role"]; exists {
					tflog.Warn(ctx, "in role")
					switch v := role.(type) {
					case string:
						if v == "#" {
							tflog.Warn(ctx, "converting")
							resultMap["role"] = 0
						}
					}
				}
			}
		}
	}

	// Normalize the "scan_allowed" field
	for _, item := range rawData {
		if scanAllowed, exists := item["scan_allowed"]; exists {
			scans, ok := scanAllowed.([]interface{})
			if !ok {
				tflog.Info(ctx, "Error processing JSON response from the server")
			}

			for i, scan := range scans {
				switch v := scan.(type) {
				case string:
					if v == "#" {
						scans[i] = 0
					}
				}
			}
		}
	}

	// Step 3: Remarshal the normalized data into the final struct
	// var items []Workflow
	normalizedJSON, err := json.Marshal(rawData)
	tflog.Warn(ctx, string(normalizedJSON))
	if err != nil {
		tflog.Error(ctx, "Error marshaling normalized data")
	}

	return normalizedJSON
}