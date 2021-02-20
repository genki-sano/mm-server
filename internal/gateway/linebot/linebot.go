package linebot

import (
	"context"
	"net/url"

	"github.com/genki-sano/mm-server/internal/gateway"
)

const (
	// APIEndpointVerifyAccessToken const
	APIEndpointVerifyAccessToken = "/oauth2/v2.1/verify"

	// APIEndpointGetProfile const
	APIEndpointGetProfile = "/v2/profile"
)

// LineGateway type
type LineGateway struct{}

// NewLineGateway method
func NewLineGateway() gateway.LineDataAccess {
	return &LineGateway{}
}

// VerifyToken method
func (g *LineGateway) VerifyToken(ctx context.Context, token string) (*gateway.VerifyTokenResponse, error) {
	client := newClient()

	vs := url.Values{}
	vs.Set("access_token", token)

	res, err := client.get(ctx, APIEndpointVerifyAccessToken, vs)
	if err != nil {
		return nil, err
	}

	return decodeToVerifyTokenResponse(res)
}

// GetProfile method
func (g *LineGateway) GetProfile(ctx context.Context, token string) (*gateway.UserProfileResponse, error) {
	client := newClient()

	vs := url.Values{}

	res, err := client.withAceesToken(token).get(ctx, APIEndpointVerifyAccessToken, vs)
	if err != nil {
		return nil, err
	}

	return decodeToUserProfileResponse(res)
}
