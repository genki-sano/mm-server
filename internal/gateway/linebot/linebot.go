package linebot

import (
	"context"
	"net/url"

	"github.com/genki-sano/mm-server/internal/gateway"
)

type lineGateway struct{}

// NewLineGateway method
func NewLineGateway() gateway.LineDataAccess {
	return &lineGateway{}
}

func (g *lineGateway) VerifyToken(ctx context.Context, token string) (*gateway.VerifyTokenResponse, error) {
	client, err := newClient()
	if err != nil {
		return nil, err
	}

	vs := url.Values{}
	vs.Set("access_token", token)

	res, err := client.get(ctx, APIEndpointVerifyToken, vs)
	if err != nil {
		return nil, err
	}

	defer closeResponse(res)
	return decodeToVerifyTokenResponse(res)
}

func (g *lineGateway) GetProfile(ctx context.Context, token string) (*gateway.UserProfileResponse, error) {
	client, err := newClient(withAceesToken(token))
	if err != nil {
		return nil, err
	}

	res, err := client.get(ctx, APIEndpointGetProfile, nil)
	if err != nil {
		return nil, err
	}

	defer closeResponse(res)
	return decodeToUserProfileResponse(res)
}
