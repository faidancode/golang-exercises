package cart

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type DummyClient struct {
	BaseURL string
	client  *http.Client //pointer client
}

func NewDummyClient() *DummyClient {
	return &DummyClient{
		BaseURL: "https://dummyjson.com",
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *DummyClient) FetchGeneric(endpoint string, target interface{}) error {
	resp, err := http.Get(c.BaseURL + endpoint)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Error API: status %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}
