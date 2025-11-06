package client

import (
	"net/http"
	"time"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Token      string
}

func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
