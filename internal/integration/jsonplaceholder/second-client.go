package jsonplaceholder

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SecondClient struct {
	BaseURL string
	client  *http.Client
}

func NewSecondClient() *SecondClient {
	return &SecondClient{
		BaseURL: "https://jsonplaceholder.typicode.com",
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *SecondClient) Get(endpoint string, target interface{}) error {
	url := c.BaseURL + endpoint
	resp, err := c.client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Request failed with status %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}
