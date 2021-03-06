package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/genki-sano/mm-server/cmd/http/helper"
	"github.com/genki-sano/mm-server/internal/usecase"
)

// ListController type
type ListController struct {
	u usecase.UserListUsecase
}

// NewListController method
func NewListController(u usecase.UserListUsecase) *ListController {
	return &ListController{
		u: u,
	}
}

// Handler method
func (c *ListController) Handler(ctx *gin.Context) {
	res, err := c.u.Handle(ctx)
	if err != nil {
		helper.Response(ctx, helper.CreateErrorResponseData(http.StatusBadRequest, err))
		return
	}

	helper.Response(ctx, helper.CreateSuccessResponseData(res))
}
