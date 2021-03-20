package usecase

import (
	"context"

	"github.com/genki-sano/mm-server/internal/presenter"
	"github.com/genki-sano/mm-server/internal/valueobject"
)

// PaytmentListUsecase type
type PaytmentListUsecase interface {
	Handle(context.Context, *valueobject.Month) (presenter.I, error)
}
