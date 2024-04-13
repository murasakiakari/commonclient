package commonclient

import (
	"net/http"
	"sync"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type SingleThreadedHttpClient struct {
	lock   sync.Mutex
	client *http.Client
}

func NewSingleThreadedHttpClient(client *http.Client) *SingleThreadedHttpClient {
	return &SingleThreadedHttpClient{
		client: client,
	}
}

func (c *SingleThreadedHttpClient) Do(req *http.Request) (*http.Response, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	return c.client.Do(req)
}
