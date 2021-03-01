package linebot

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

// APIEndpoint constants
const (
	APIEndpointBase = "https://api.line.me"

	APIEndpointVerifyToken = "/oauth2/v2.1/verify"
	APIEndpointGetProfile  = "/v2/profile"
)

type client struct {
	endpointBase *url.URL
	httpClient   *http.Client
	accessToken  string
}

type clientOption func(*client) error

func newClient(options ...clientOption) (*client, error) {
	c := &client{
		httpClient: http.DefaultClient,
	}
	u, err := url.ParseRequestURI(APIEndpointBase)
	if err != nil {
		return nil, err
	}
	c.endpointBase = u

	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func withAceesToken(token string) clientOption {
	return func(c *client) error {
		c.accessToken = token
		return nil
	}
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
	req, err := http.NewRequest(http.MethodGet, c.url(c.endpointBase, endpoint), nil)
	if err != nil {
		return nil, err
	}
	if query != nil {
		req.URL.RawQuery = query.Encode()
	}
	return c.do(ctx, req)
}

func (c *client) post(ctx context.Context, endpoint string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, c.url(c.endpointBase, endpoint), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return c.do(ctx, req)
}

// @see https://github.com/google/go-github/pull/317
func closeResponse(res *http.Response) error {
	defer res.Body.Close()
	_, err := io.Copy(ioutil.Discard, res.Body)
	return err
}
