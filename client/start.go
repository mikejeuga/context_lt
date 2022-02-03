package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	ctx := context.Background()
	ctxT, cancel := context.WithTimeout(ctx, 12 * time.Second)
	defer cancel()

	client := NewClient(http.DefaultClient)
	client.Get(ctxT, "http://localhost:8094")

	client.Get(ctxT, "http://localhost:8094/far")

}

type Client struct {
	client *http.Client
}

func NewClient(client *http.Client) *Client {
	return &Client{client: client}
}


func (c *Client) Get(ctx context.Context, url string) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}
	io.Copy(os.Stdout, res.Body)
}

