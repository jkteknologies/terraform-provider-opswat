package opswatClient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// GetQueue - Returns global sync scan queue
func (c *Client) GetQueue(ctx context.Context) (*Queue, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/admin/config/scan", c.HostURL), nil)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "request URL: "+fmt.Sprintf("%s/admin/config/scan", c.HostURL))
	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := Queue{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateQueue - Update global sync scan queue
func (c *Client) UpdateQueue(ctx context.Context, queue int) (*Queue, error) {
	queueJson := map[string]int{"max_queue_per_agent": queue}
	preparedJson, err := json.Marshal(queueJson)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/config/scan", c.HostURL), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "request URL: "+fmt.Sprintf("%s/admin/config/scan, request body: %s", c.HostURL, string(preparedJson)))
	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := Queue{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateQueue - Creates global sync scan queue
func (c *Client) CreateQueue(ctx context.Context, queue int) (*Queue, error) {
	queueJson := map[string]int{"max_queue_per_agent": queue}
	preparedJson, err := json.Marshal(queueJson)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/config/scan", c.HostURL), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "request URL: "+fmt.Sprintf("%s/admin/config/scan, request body: %s", c.HostURL, string(preparedJson)))
	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	result := Queue{}

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
