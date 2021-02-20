package firebase

import (
	"context"

	fb "firebase.google.com/go/v4"

	"github.com/genki-sano/mm-server/internal/gateway"
)

// FirebaseGateway type
type FirebaseGateway struct{}

// NewFirebaseGateway method
func NewFirebaseGateway() gateway.FirebaseDataAccess {
	return &FirebaseGateway{}
}

// CreateCustomToken method
func (g *FirebaseGateway) CreateCustomToken(ctx context.Context, uid string) (string, error) {
	app, err := fb.NewApp(context.Background(), nil)
	if err != nil {
		return "", err
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		return "", err
	}

	token, err := client.CustomToken(ctx, uid)
	if err != nil {
		return "", err
	}

	return token, nil
}
