// +build wireinject

package di

import (
	"github.com/google/wire"

	controller "github.com/genki-sano/mm-server/cmd/http/controller/payment"
	"github.com/genki-sano/mm-server/internal/gateway/linebot"
	"github.com/genki-sano/mm-server/internal/gateway/spreadsheet"
	presenter "github.com/genki-sano/mm-server/internal/presenter/payment"
	usecase "github.com/genki-sano/mm-server/internal/usecase/payment"
)

func InitializePaymentCreate() *controller.CreateController {
	wire.Build(
		controller.NewCreateController,
		spreadsheet.NewUserRepository,
		spreadsheet.NewPaymentRepository,
		linebot.NewLineGateway,
		usecase.NewCreateUsecase,
	)
	return nil
}

func InitializePaymentList() *controller.ListController {
	wire.Build(
		controller.NewListController,
		spreadsheet.NewPaymentRepository,
		presenter.NewListFactory,
		usecase.NewListUsecase,
	)
	return nil
}
