package linebot

import (
	"context"
	"net/url"
	"os"

	"github.com/genki-sano/mm-server/internal/gateway"
)

type lineGateway struct {
	secret string
	token  string
}

// NewLineGateway method
func NewLineGateway() gateway.LineDataAccess {
	return &lineGateway{
		secret: os.Getenv("LINE_MESSAGE_CHANNEL_SECRET"),
		token:  os.Getenv("LINE_MESSAGE_CHANNEL_ACCESS_TOKEN"),
	}
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

func (g *lineGateway) PushFlexMessage(to, altText string, data []byte) error {
	clinet, err := newMessagingApi(g.secret, g.token)
	if err != nil {
		return err
	}

	message, err := newFlexMessage(altText, data).do()
	if err != nil {
		return err
	}

	if err := clinet.newPushMessage(to, message).do(); err != nil {
		return err
	}
	return nil
}
