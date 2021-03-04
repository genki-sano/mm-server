package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/genki-sano/mm-server/cmd/http/helper"
	"github.com/genki-sano/mm-server/internal/usecase"
	"github.com/genki-sano/mm-server/internal/valueobject"
)

// ListController type
type ListController struct {
	u usecase.PaytmentListUsecase
}

// NewListController method
func NewListController(u usecase.PaytmentListUsecase) *ListController {
	return &ListController{
		u: u,
	}
}

// Handler method
func (c *ListController) Handler(ctx *gin.Context) {
	date, err := valueobject.NewPaymentDate(ctx.Query("date"))
	if err != nil {
		helper.Response(ctx, helper.CreateErrorResponseData(http.StatusBadRequest, err))
		return
	}

	res, err := c.u.Handle(ctx, date)
	if err != nil {
		helper.Response(ctx, helper.CreateErrorResponseData(http.StatusBadRequest, err))
		return
	}

	helper.Response(ctx, helper.CreateSuccessResponseData(res))
}
