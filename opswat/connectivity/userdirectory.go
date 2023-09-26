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

// GetDirs - Returns global sync scan timeout
func (c *Client) GetDirs() ([]UserDirectory, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/admin/userdirectory", c.HostURL), nil)

	if err != nil {
		return nil, err
	}

	fmt.Println("request URL: " + fmt.Sprintf("%s/admin/userdirectory", c.HostURL))

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	var result []UserDirectory

	err = json.Unmarshal(body, &result)

	ctx := context.TODO()
	justString := fmt.Sprint(&result)
	tflog.Info(ctx, justString)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetDir - Returns global sync scan timeout
func (c *Client) GetDir(dirId int) (*UserDirectory, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/admin/userdirectory", c.HostURL), nil)

	if err != nil {
		return nil, err
	}

	fmt.Println("request URL: " + fmt.Sprintf("%s/admin/userdirectory/%d", c.HostURL, dirId))

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := UserDirectory{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateDir - Update global sync scan timeout
func (c *Client) UpdateDir(dirId int, userDir UserDirectory) (*UserDirectory, error) {

	preparedJson, err := json.Marshal(userDir)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/userdirectory/%d", c.HostURL, dirId), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := UserDirectory{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateDir - Creates global sync scan timeout
func (c *Client) CreateDir(userDir UserDirectory) (*UserDirectory, error) {
	preparedJson, err := json.Marshal(userDir)
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

	result := UserDirectory{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteUserDirectory - Delete userdirectory
func (c *Client) DeleteDir(dirID int) error {

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
