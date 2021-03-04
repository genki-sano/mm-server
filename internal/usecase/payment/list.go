package payment

import (
	"context"

	"github.com/genki-sano/mm-server/internal/gateway"
	"github.com/genki-sano/mm-server/internal/presenter"
	"github.com/genki-sano/mm-server/internal/usecase"
	"github.com/genki-sano/mm-server/internal/valueobject"
)

type listInteractor struct {
	paymentRepos gateway.PaymentDataAccess
	presenter    presenter.PaymentListFactory
}

// NewListUsecase method
func NewListUsecase(
	paymentRepos gateway.PaymentDataAccess,
	presenter presenter.PaymentListFactory,
) usecase.PaytmentListUsecase {
	return &listInteractor{
		paymentRepos: paymentRepos,
		presenter:    presenter,
	}
}

// Handle method
func (i *listInteractor) Handle(
	ctx context.Context,
	date *valueobject.PaymentDate,
) (presenter.I, error) {
	payments, err := i.paymentRepos.GetByDate(date.Get())
	if err != nil {
		return nil, err
	}

	return i.presenter.New(payments), nil
}
