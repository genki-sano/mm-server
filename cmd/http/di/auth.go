// +build wireinject

package di

import (
	"github.com/google/wire"

	controller "github.com/genki-sano/mm-server/cmd/http/controller/auth"
	"github.com/genki-sano/mm-server/internal/gateway/firebase"
	"github.com/genki-sano/mm-server/internal/gateway/linebot"
	"github.com/genki-sano/mm-server/internal/gateway/spreadsheet"
	presenter "github.com/genki-sano/mm-server/internal/presenter/auth"
	usecase "github.com/genki-sano/mm-server/internal/usecase/auth"
)

func InitializeAuthVerify() *controller.VerifyController {
	wire.Build(
		controller.NewVerifyController,
		firebase.NewFirebaseGateway,
		linebot.NewLineGateway,
		spreadsheet.NewUserRepository,
		presenter.NewVerifyFactory,
		usecase.NewVerifyUsecase,
	)
	return nil
}
