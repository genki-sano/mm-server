// +build wireinject

package di

import (
	"github.com/google/wire"

	controller "github.com/genki-sano/mm-server/cmd/http/controller/user"
	"github.com/genki-sano/mm-server/internal/gateway/spreadsheet"
	presenter "github.com/genki-sano/mm-server/internal/presenter/user"
	usecase "github.com/genki-sano/mm-server/internal/usecase/user"
)

func InitializeUserList() *controller.ListController {
	wire.Build(
		controller.NewListController,
		spreadsheet.NewUserRepository,
		presenter.NewListFactory,
		usecase.NewListUsecase,
	)
	return nil
}
