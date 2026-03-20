package jsonplaceholder

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Kenapa pakai *http.Client? bisa reuse, timeout.
type Client struct {
	BaseURL string
	client  *http.Client
}

// * (Asterisk) adalah operator untuk Mendefinisikan tipe pointer atau Mengintip isinya (dereference).
// & (Ampersand) adalah operator untuk Menciptakan pointer (mengambil alamat).

// Function ini adalah constructor untuk membuat instance Client.
// Kenapa return pointer *Client ? Tidak meng-copy struct, efisien memory, bisa dipakai oleh handler/service.
func NewClient() *Client {
	// &Client berarti: ambil alamat memory dari struct Client
	// Jadi yang dikembalikan adalah: client -> memory address -> Client struct
	return &Client{
		BaseURL: "https://jsonplaceholder.typicode.com",
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
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

	// Deferred Function Call.
	// Memastikan sebuah perintah tetap dijalankan di akhir fungsi, apa pun yang terjadi.
	// Tanpa defer, berarti harus menulis resp.Body.Close() secara manual sebelum setiap baris return.
	// .Close mencegah memory leaks
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status %d", resp.StatusCode)
	}

	// Decoder membutuhkan pointer ke variable agar bisa menulis data.
	// Jadi: &posts berarti: alamat memory posts
	return json.NewDecoder(resp.Body).Decode(target)
}
