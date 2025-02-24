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

// GetUser - Returns User by id
func (c *Client) GetUser(ctx context.Context, userID int) (*User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/admin/user/%d", c.HostURL, userID), nil)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "request URL: " + fmt.Sprintf("%s/admin/user/%d", c.HostURL, userID))
	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := User{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateUser - Update User by id
func (c *Client) UpdateUser(ctx context.Context, userID int, config User) (*User, error) {

	preparedJson, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/user/%d", c.HostURL, userID), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "request URL: " + fmt.Sprintf("%s/admin/user/%d, request body: %s", c.HostURL, userID, string(preparedJson)))
	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := User{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateUser - Creates User
func (c *Client) CreateUser(ctx context.Context, config User) (*User, error) {

	preparedJson, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/admin/user", c.HostURL), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "request URL: " + fmt.Sprintf("%s/admin/user/%s", c.HostURL, string(preparedJson)))
	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := User{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteUser - Delete user by id
func (c *Client) DeleteUser(ctx context.Context, userID int) error {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/admin/user/%d", c.HostURL, userID), nil)
	if err != nil {
		return err
	}

	tflog.Debug(ctx, "request URL: " + fmt.Sprintf("%s/admin/user/%d", c.HostURL, userID))
	body, err := c.doRequest(req)

	if err != nil {
		return err
	}

	if string(body) != `{"result":"Success"}` {
		return errors.New(string(body))
	}

	return nil
}
