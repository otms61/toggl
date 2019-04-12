package toggl

import (
	"log"
	"net/http"
	"os"
)

// APIURL added as a var so that we can change this for testing purposes
var APIURL = "https://www.toggl.com/api/"

// Option defines an option for a Client
type Option func(*Client)

// OptionDebug enable debugging for the client.
func OptionDebug(b bool) func(*Client) {
	return func(c *Client) {
		c.debug = b
	}
}

// OptionHTTPClient enable to use a custom HTTPClient.
func OptionHTTPClient(hc *http.Client) func(*Client) {
	return func(c *Client) {
		c.HTTPClient = hc
	}
}

// New builds a toggl client from the provided token.
func New(token string, options ...Option) *Client {
	t := &Client{
		token:      token,
		HTTPClient: &http.Client{},
		logger:     log.New(os.Stderr, "otms61/toggl", log.LstdFlags|log.Lshortfile),
	}

	for _, opt := range options {
		opt(t)
	}

	return t
}
