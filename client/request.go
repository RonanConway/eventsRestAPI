package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RonanConway/eventsRestAPI/models"
)

func (c *Client) doRequest(method, path string, body any, auth bool) (*http.Response, error) {

	url := fmt.Sprintf("%s%s", c.BaseURL, path)

	var reqBody []byte
	var err error

	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if auth && c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}

	return c.HTTPClient.Do(req)
}

func (c *Client) GetEvents() ([]models.Event, error) {
	resp, err := c.doRequest(http.MethodGet, "/events", nil, false)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var events []models.Event
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, err
	}
	return events, nil
}

func (c *Client) CreateEvent(event models.Event) error {
	_, err := c.doRequest(http.MethodPost, "/events", event, true)
	return err
}
