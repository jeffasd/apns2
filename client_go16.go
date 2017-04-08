// +build go1.6,!go1.7

package apns2

import (
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
)

func (c *Client) requestWithContext(ctx context.Context, req *http.Request) (*http.Response, error) {
	var httpRes *http.Response
	if ctx != nil {
		return ctxhttp.Do(ctx, c.HTTPClient, req)
	}
	return c.HTTPClient.Do(req)
}
