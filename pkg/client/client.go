package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/cookiejar"
	"time"
)

// Client gets things
type Client interface {
	GET(url string) (*bytes.Buffer, error)
}

// Option allows for modifying a httpClient
type Option func(*httpClient)

// SetClient sets the underlying http.Client
func SetClient(client *http.Client) Option {
	return func(c *httpClient) {
		c.client = client
	}
}

type httpClient struct {
	client *http.Client
}

// New creates a new Client
func New(opts ...Option) Client {
	jar, _ := cookiejar.New(nil)
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          20,
		IdleConnTimeout:       30 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	c := httpClient{
		client: &http.Client{
			Jar:       jar,
			Transport: transport,
			Timeout:   30 * time.Second,
		},
	}

	for i := range opts {
		opts[i](&c)
	}

	return &c
}

// GET makes a get request to the provided url. It will read the response body fully and return
// a copy of those bytes so the caller does not have to manage the connection.
func (c *httpClient) GET(url string) (*bytes.Buffer, error) {
	//resp, err := http.Get(url)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("record not found")
		}
		return nil, fmt.Errorf("error making request to server")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(b), nil
}
