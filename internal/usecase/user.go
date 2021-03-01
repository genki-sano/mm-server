package usecase

import (
	"github.com/genki-sano/mm-server/internal/presenter"
)

// UserListUsecase interface
type UserListUsecase interface {
	Handle() (presenter.I, error)
}
