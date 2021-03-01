// +build wireinject

package di

import (
	"github.com/google/wire"

	controller "github.com/genki-sano/mm-server/cmd/http/controller/user"
	repository "github.com/genki-sano/mm-server/internal/gateway/database"
	infarastructure "github.com/genki-sano/mm-server/internal/infarastructure/database"
	presenter "github.com/genki-sano/mm-server/internal/presenter/user"
	usecase "github.com/genki-sano/mm-server/internal/usecase/user"
)

func InitializeUserList() *controller.ListController {
	wire.Build(
		controller.NewListController,
		usecase.NewListInteractor,
		presenter.NewListFactory,
		repository.NewUserRepository,
		infarastructure.NewSQLHandler,
	)
	return nil
}
