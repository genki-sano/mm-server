package payment

import (
	"context"

	"github.com/genki-sano/mm-server/internal/entity"
	"github.com/genki-sano/mm-server/internal/gateway"
	"github.com/genki-sano/mm-server/internal/usecase"
	"github.com/genki-sano/mm-server/internal/valueobject"
)

type createInteractor struct {
	paymentRepos gateway.PaymentDataAccess
}

// NewCreateUsecase method
func NewCreateUsecase(
	paymentRepos gateway.PaymentDataAccess,
) usecase.PaytmentCreateUsecase {
	return &createInteractor{
		paymentRepos: paymentRepos,
	}
}

// Handle method
func (i *createInteractor) Handle(
	ctx context.Context,
	userType *valueobject.UserType,
	category *valueobject.Category,
	price *valueobject.Price,
	date *valueobject.Date,
	memo *valueobject.Memo,
) error {
	payment := &entity.Payment{
		ID:       0, // 仮の値を入れておく
		UserType: userType.Get(),
		Category: category.Get(),
		Price:    price.Get(),
		Date:     date.Get(),
		Memo:     memo.Get(),
	}
	if err := i.paymentRepos.Insert(payment); err != nil {
		return err
	}

	return nil
}
