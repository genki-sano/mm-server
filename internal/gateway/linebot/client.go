package linebot

import (
	"context"
	"net/http"
	"net/url"
	"path"
)

// APIEndpointBase const
const APIEndpointBase = "https://api.line.me"

type client struct {
	accessToken string
	httpClient  *http.Client
}

func newClient() *client {
	return &client{
		accessToken: "",
		httpClient:  http.DefaultClient,
	}
}

func (c *client) withAceesToken(token string) *client {
	c.accessToken = token
	return c
}

func (c *client) url(base *url.URL, endpoint string) string {
	u := *base
	u.Path = path.Join(u.Path, endpoint)
	return u.String()
}

func (c *client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	if c.accessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.accessToken)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	return c.httpClient.Do(req)

}

func (c *client) get(ctx context.Context, endpoint string, query url.Values) (*http.Response, error) {
	u, err := url.ParseRequestURI(APIEndpointBase)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, c.url(u, endpoint), nil)
	if err != nil {
		return nil, err
	}
	if query != nil {
		req.URL.RawQuery = query.Encode()
	}
	return c.do(ctx, req)
}
