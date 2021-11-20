// Package artipie provides API for managing artipie server.
package artipie

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

// A Client manages communication with the Artipie API.
type Client struct {
	client    *http.Client
	Endpoint  *url.URL
	UserAgent string
	Auth      Auth
}

// NewClient returns a new Artipie API client. If a nil httpClient is
// provided, a new http.Client will be used.
// auth should implement Auth interface and be able to set authentification headers to request.
// auth could be nil, in that case http.client should provide authentification.
func NewClient(client *http.Client, endpoint string, auth Auth) (*Client, error) {
	if client == nil {
		client = &http.Client{}
	}
	if endpoint == "" {
		return nil, errors.New("Endpoint is not specified")
	}
	ep, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	return &Client{client: client, Endpoint: ep, Auth: auth}, nil
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	u, err := c.Endpoint.Parse(path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	if c.Auth != nil {
		c.Auth.SetAuthHeader(req)
	}
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// NewJSONRequest creates an API request with JSON body.
// If specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewJSONRequest(method, path string, body interface{}) (*http.Request, error) {
	if body == nil {
		return c.NewRequest(method, path, nil)
	}
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(body)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest(method, path, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	return req, nil
}
