package user

import (
	"github.com/genki-sano/mm-server/internal/gateway/database"
	"github.com/genki-sano/mm-server/internal/presenter"
	"github.com/genki-sano/mm-server/internal/usecase"
)

type listInteractor struct {
	repository database.UserDataAccess
	presenter  presenter.UserListFactory
}

// NewListInteractor method
func NewListInteractor(
	repository database.UserDataAccess,
	presenter presenter.UserListFactory,
) usecase.UserListUsecase {
	return &listInteractor{
		repository: repository,
		presenter:  presenter,
	}
}

// Handle method
func (i *listInteractor) Handle() (presenter.I, error) {
	users, err := i.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return i.presenter.New(users), nil
}
