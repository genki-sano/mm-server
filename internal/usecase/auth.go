package usecase

import (
	"context"

	"github.com/genki-sano/mm-server/internal/presenter"
	"github.com/genki-sano/mm-server/internal/valueobject"
)

// AuthVerifyUsecase type
type AuthVerifyUsecase interface {
	Handle(context.Context, *valueobject.AccessToken) (presenter.I, error)
}
