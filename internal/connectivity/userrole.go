package opswatClient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/TylerBrock/colorjson"
	"github.com/emirpasic/gods/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"net/http"
	"strings"
)

// GetUserRole - Returns User role by id
func (c *Client) GetUserRole(userID int) (*UserRole, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/admin/role/%d", c.HostURL, userID), nil)

	if err != nil {
		return nil, err
	}

	fmt.Println("request URL: " + fmt.Sprintf("%s/admin/role/%d", c.HostURL, userID))

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
func (c *Client) UpdateUserRole(userID int, config UserRole) (*UserRole, error) {

	preparedJson, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	fmt.Println("----------- REQUEST -------------")
	fmt.Println(string(preparedJson), err)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/role/%d", c.HostURL, userID), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

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
func (c *Client) CreateUserRole(config UserRole) (*UserRole, error) {

	preparedJson, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/admin/role", c.HostURL), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	tflog.Info(ctx, utils.ToString(preparedJson))

	f := colorjson.NewFormatter()
	f.Indent = 4
	fmt.Println(string(preparedJson))

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
func (c *Client) DeleteUserRole(roleID int) error {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/admin/role/%d", c.HostURL, roleID), nil)
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
