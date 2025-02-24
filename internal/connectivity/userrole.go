package opswatClient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"net/http"
	"strings"
)

// GetUserRole - Returns User role by id
func (c *Client) GetUserRole(ctx context.Context, userID int) (*UserRole, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/admin/role/%d", c.HostURL, userID), nil)
	
	if err != nil {
		return nil, err
	}
	
	tflog.Debug(ctx, "request URL: " + fmt.Sprintf("%s/admin/role%d", c.HostURL, userID))
	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := UserRole{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateUserRole - Update User role by id
func (c *Client) UpdateUserRole(ctx context.Context, userID int, config UserRole) (*UserRole, error) {
	preparedJson, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/role/%d", c.HostURL, userID), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "request URL: " + fmt.Sprintf("%s/admin/role/%d, request body: %s", c.HostURL, userID, string(preparedJson)))
	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := UserRole{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateUser - Creates User
func (c *Client) CreateUserRole(ctx context.Context, config UserRole) (*UserRole, error) {

	preparedJson, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/admin/role", c.HostURL), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "request URL: " + fmt.Sprintf("%s/admin/role, request body: %s", c.HostURL, string(preparedJson)))
	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := UserRole{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteUserRole - Delete user role by id
func (c *Client) DeleteUserRole(ctx context.Context, roleID int) error {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/admin/role/%d", c.HostURL, roleID), nil)
	if err != nil {
		return err
	}

	tflog.Debug(ctx, "request URL: " + fmt.Sprintf("%s/admin/role/%d", c.HostURL, roleID))
	body, err := c.doRequest(req)

	if err != nil {
		return err
	}

	if string(body) != `{"result":"Success"}` {
		return errors.New(string(body))
	}

	return nil
}
