package auth

import (
	"context"
	"errors"
	"os"

	"github.com/genki-sano/mm-server/internal/gateway"
	"github.com/genki-sano/mm-server/internal/presenter"
	"github.com/genki-sano/mm-server/internal/usecase"
	"github.com/genki-sano/mm-server/internal/valueobject"
)

type verifyInteractor struct {
	firebase  gateway.FirebaseDataAccess
	linebot   gateway.LineDataAccess
	presenter presenter.AuthVerifyFactory
}

// NewVerifyUsecase method
func NewVerifyUsecase(
	firebase gateway.FirebaseDataAccess,
	linebot gateway.LineDataAccess,
	presenter presenter.AuthVerifyFactory,
) usecase.AuthVerifyUsecase {
	return &verifyInteractor{
		linebot:   linebot,
		firebase:  firebase,
		presenter: presenter,
	}
}

// Handle method
func (i *verifyInteractor) Handle(
	ctx context.Context,
	token *valueobject.AccessToken,
) (presenter.I, error) {
	tk, err := i.linebot.VerifyToken(ctx, token.Get())
	if err != nil {
		return nil, err
	}
	if tk.ClientID != os.Getenv("LINE_LOGIN_CHANNEL_ID") {
		return nil, errors.New("アクセストークンが不正です。")
	}
	if tk.ExpiresIn > 0 {
		return nil, errors.New("アクセストークンの期限が切れています。")
	}

	pf, err := i.linebot.GetProfile(ctx, token.Get())
	if err != nil {
		return nil, err
	}

	ct, err := i.firebase.CreateCustomToken(ctx, pf.UserID)
	if err != nil {
		return nil, err
	}
	return i.presenter.New(ct), nil
}
