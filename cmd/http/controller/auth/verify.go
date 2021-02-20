package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/genki-sano/mm-server/cmd/http/helper"
	"github.com/genki-sano/mm-server/internal/usecase"
	"github.com/genki-sano/mm-server/internal/valueobject"
)

// VerifyController type
type VerifyController struct {
	u usecase.AuthVerifyUsecase
}

// NewVerifyController method
func NewVerifyController(u usecase.AuthVerifyUsecase) *VerifyController {
	return &VerifyController{
		u: u,
	}
}

// Handler method
func (c *VerifyController) Handler(ctx *gin.Context) {
	accessToken, err := valueobject.NewAccessToken(ctx.Query("token"))
	if err != nil {
		helper.Response(ctx, helper.CreateErrorResponseData(http.StatusInternalServerError, err))
		return
	}

	res, err := c.u.Handle(ctx, accessToken)
	if err != nil {
		helper.Response(ctx, helper.CreateErrorResponseData(http.StatusInternalServerError, err))
		return
	}

	helper.Response(ctx, helper.CreateSuccessResponseData(res))
}
