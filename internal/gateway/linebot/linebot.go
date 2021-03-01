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

type lineGateway struct{}

// NewLineGateway method
func NewLineGateway() gateway.LineDataAccess {
	return &lineGateway{}
}

func (g *lineGateway) VerifyToken(ctx context.Context, token string) (*gateway.VerifyTokenResponse, error) {
	client := newClient()

	vs := url.Values{}
	vs.Set("access_token", token)

	res, err := client.get(ctx, APIEndpointVerifyAccessToken, vs)
	if err != nil {
		return nil, err
	}

	return decodeToVerifyTokenResponse(res)
}

func (g *lineGateway) GetProfile(ctx context.Context, token string) (*gateway.UserProfileResponse, error) {
	client := newClient()

	res, err := client.withAceesToken(token).get(ctx, APIEndpointGetProfile, nil)
	if err != nil {
		return nil, err
	}

	return decodeToUserProfileResponse(res)
}
