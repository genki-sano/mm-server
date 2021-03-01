package firebase

import (
	"context"

	fb "firebase.google.com/go/v4"

	"github.com/genki-sano/mm-server/internal/gateway"
)

type firebaseGateway struct{}

// NewFirebaseGateway method
func NewFirebaseGateway() gateway.FirebaseDataAccess {
	return &firebaseGateway{}
}

func (g *firebaseGateway) CreateCustomToken(ctx context.Context, uid string) (string, error) {
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
