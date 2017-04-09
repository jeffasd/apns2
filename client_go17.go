// +build go1.7

package apns2

import "net/http"

func (c *Client) requestWithContext(ctx context, req *http.Request) (*http.Response, error) {
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	return c.HTTPClient.Do(req)
}
