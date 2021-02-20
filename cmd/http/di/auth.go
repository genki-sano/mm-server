// +build wireinject

package di

import (
	"github.com/google/wire"

	controller "github.com/genki-sano/mm-server/cmd/http/controller/auth"
	"github.com/genki-sano/mm-server/internal/gateway/firebase"
	"github.com/genki-sano/mm-server/internal/gateway/linebot"
	presenter "github.com/genki-sano/mm-server/internal/presenter/auth"
	usecase "github.com/genki-sano/mm-server/internal/usecase/auth"
)

func InitializeAuthVerify() *controller.VerifyController {
	wire.Build(
		controller.NewVerifyController,
		linebot.NewLineGateway,
		firebase.NewFirebaseGateway,
		presenter.NewVerifyFactory,
		usecase.NewVerifyUsecase,
	)
	return nil
}
