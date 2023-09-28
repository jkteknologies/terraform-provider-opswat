package opswatClient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetQueue - Returns global sync scan queue
func (c *Client) GetQueue() (*Queue, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/admin/config/scan", c.HostURL), nil)

	if err != nil {
		return nil, err
	}

	fmt.Println("request URL: " + fmt.Sprintf("%s/admin/config/scan", c.HostURL))

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
func (c *Client) UpdateQueue(queue int) (*Queue, error) {
	queueJson := map[string]int{"max_queue_per_agent": queue}
	preparedJson, err := json.Marshal(queueJson)
	if err != nil {
		return nil, err
	}

	fmt.Println("----------- REQUEST -------------")
	fmt.Println(string(preparedJson), err)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/config/scan", c.HostURL), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

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
func (c *Client) CreateQueue(queue int) (*Queue, error) {
	queueJson := map[string]int{"max_queue_per_agent": queue}
	preparedJson, err := json.Marshal(queueJson)
	if err != nil {
		return nil, err
	}

	fmt.Println("----------- REQUEST -------------")
	fmt.Println(string(preparedJson), err)

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/admin/config/scan", c.HostURL), strings.NewReader(string(preparedJson)))
	if err != nil {
		return nil, err
	}

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
