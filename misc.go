package toggl

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

func (c *Client) parseResponseBody(body io.ReadCloser, intf interface{}) error {
	response, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	if c.Debug() {
		c.Debugln("parseResponseBody", string(response))
	}

	return json.Unmarshal(response, intf)
}

func (c *Client) logResponse(resp *http.Response) error {
	if c.Debug() {
		text, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return err
		}
		c.Debugln(string(text))
	}
	return nil
}
