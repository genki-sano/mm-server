package user

import (
	"context"

	"github.com/genki-sano/mm-server/internal/gateway"
	"github.com/genki-sano/mm-server/internal/presenter"
	"github.com/genki-sano/mm-server/internal/usecase"
)

type listInteractor struct {
	userRepos gateway.UserDataAccess
	presenter presenter.UserListFactory
}

// NewListUsecase method
func NewListUsecase(
	userRepos gateway.UserDataAccess,
	presenter presenter.UserListFactory,
) usecase.UserListUsecase {
	return &listInteractor{
		userRepos: userRepos,
		presenter: presenter,
	}
}

// Handle method
func (i *listInteractor) Handle(ctx context.Context) (presenter.I, error) {
	payments, err := i.userRepos.GetAll()
	if err != nil {
		return nil, err
	}

	return i.presenter.New(payments), nil
}
