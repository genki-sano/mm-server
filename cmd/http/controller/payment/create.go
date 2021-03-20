package payment

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/genki-sano/mm-server/cmd/http/helper"
	"github.com/genki-sano/mm-server/internal/usecase"
	"github.com/genki-sano/mm-server/internal/valueobject"
)

// CreateController type
type CreateController struct {
	u usecase.PaytmentCreateUsecase
}

// NewCreateController method
func NewCreateController(u usecase.PaytmentCreateUsecase) *CreateController {
	return &CreateController{
		u: u,
	}
}

// Handler method
func (c *CreateController) Handler(ctx *gin.Context) {
	userType, err := valueobject.NewUserType(ctx.PostForm("userType"))
	if err != nil {
		helper.Response(ctx, helper.CreateErrorResponseData(http.StatusBadRequest, err))
		return
	}
	category, err := valueobject.NewCategory(ctx.PostForm("category"))
	if err != nil {
		helper.Response(ctx, helper.CreateErrorResponseData(http.StatusBadRequest, err))
		return
	}
	price, err := valueobject.NewPrice(ctx.PostForm("price"))
	if err != nil {
		helper.Response(ctx, helper.CreateErrorResponseData(http.StatusBadRequest, err))
		return
	}
	date, err := valueobject.NewDate(ctx.PostForm("date"))
	if err != nil {
		helper.Response(ctx, helper.CreateErrorResponseData(http.StatusBadRequest, err))
		return
	}
	memo, err := valueobject.NewMemo(ctx.PostForm("memo"))
	if err != nil {
		helper.Response(ctx, helper.CreateErrorResponseData(http.StatusBadRequest, err))
		return
	}

	if err := c.u.Handle(ctx, userType, category, price, date, memo); err != nil {
		helper.Response(ctx, helper.CreateErrorResponseData(http.StatusBadRequest, err))
		return
	}

	helper.Response(ctx, helper.CreateSuccessNoContentResponseData())
}
