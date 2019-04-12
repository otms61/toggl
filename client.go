package toggl

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// statusCodeError represents an http response error.
type statusCodeError struct {
	Code   int
	Status string
}

func (t statusCodeError) Error() string {
	return fmt.Sprintf("Toggl server error: %s.", t.Status)
}

// Client for the toggl api.
type Client struct {
	HTTPClient *http.Client
	token      string
	logger     *log.Logger
	debug      bool
}

func (c *Client) newRequest(ctx context.Context, method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, APIURL+path, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	req.SetBasicAuth(c.token, "api_token")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "GoClient")

	return req, nil
}

func (c *Client) get(ctx context.Context, path string, values url.Values, intf interface{}) error {
	req, err := c.newRequest(ctx, "GET", path, nil)
	if err != nil {
		return err
	}

	if values != nil {
		req.URL.RawQuery = values.Encode()
	}

	return c.doRequest(ctx, req, intf)
}

func (c *Client) post(ctx context.Context, path string, body io.Reader, intf interface{}) error {
	req, err := c.newRequest(ctx, "POST", path, body)
	if err != nil {
		return err
	}

	return c.doRequest(ctx, req, intf)
}

func (c *Client) put(ctx context.Context, path string, body io.Reader, intf interface{}) error {
	req, err := c.newRequest(ctx, "PUT", path, body)
	if err != nil {
		return err
	}

	return c.doRequest(ctx, req, intf)
}

func (c *Client) delete(ctx context.Context, path string, body io.Reader, intf interface{}) error {
	req, err := c.newRequest(ctx, "DELETE", path, body)
	if err != nil {
		return err
	}

	return c.doRequest(ctx, req, intf)
}

func (c *Client) doRequest(ctx context.Context, req *http.Request, intf interface{}) error {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = c.checkStatusCode(resp)
	if err != nil {
		return err
	}

	return c.parseResponseBody(resp.Body, intf)
}

func (c *Client) checkStatusCode(resp *http.Response) error {
	if resp.StatusCode != http.StatusOK {
		return statusCodeError{Code: resp.StatusCode, Status: resp.Status}
	}

	return nil
}

// Debugf print a formatted debug line.
func (c *Client) Debugf(format string, v ...interface{}) {
	if c.debug {
		c.logger.Output(2, fmt.Sprintf(format, v...))
	}
}

// Debugln print a debug line.
func (c *Client) Debugln(v ...interface{}) {
	if c.debug {
		c.logger.Output(2, fmt.Sprintln(v...))
	}
}

// Debug returns if debug is enabled.
func (c *Client) Debug() bool {
	return c.debug
}
