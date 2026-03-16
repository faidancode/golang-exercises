package jsonplaceholder

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Kenapa pakai *http.Client? bisa reuse, timeout.
type Client struct {
	BaseURL string
	client  *http.Client
}

// Function ini adalah constructor untuk membuat instance Client.
// Kenapa return pointer *Client ? Tidak meng-copy struct, efisien memory, bisa dipakai oleh handler/service.
func NewClient() *Client {
	// &Client berarti: ambil alamat memory dari struct Client
	// Jadi yang dikembalikan adalah: client -> memory address -> Client struct
	return &Client{
		BaseURL: "https://jsonplaceholder.typicode.com",
		client:  &http.Client{},
	}
}

// (c *Client) method receiver
// target interface{} untuk menampung hasil decode JSON.
func (c *Client) Get(endpoint string, target interface{}) error {
	url := c.BaseURL + endpoint

	resp, err := c.client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status %d", resp.StatusCode)
	}

	// Decoder membutuhkan pointer ke variable agar bisa menulis data.
	// Jadi: &posts berarti: alamat memory posts
	return json.NewDecoder(resp.Body).Decode(target)
}
