// +build go1.7

package apns2

import (
	"context"
	"net/http"
)

func (c *Client) requestWithContext(ctx context.Context, req *http.Request) (*http.Response, error) {
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	return c.HTTPClient.Do(req)
}
