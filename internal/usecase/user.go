package usecase

import (
	"context"

	"github.com/genki-sano/mm-server/internal/presenter"
)

// UserListUsecase type
type UserListUsecase interface {
	Handle(context.Context) (presenter.I, error)
}
