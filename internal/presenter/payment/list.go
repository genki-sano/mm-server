package auth

import (
	"encoding/json"

	"github.com/genki-sano/mm-server/internal/entity"
	"github.com/genki-sano/mm-server/internal/presenter"
)

type listFactory struct{}

// NewListFactory method
func NewListFactory() presenter.PaymentListFactory {
	return &listFactory{}
}

type listPresenter struct {
	payments []*entity.Payment
}

// New method
func (f *listFactory) New(payments []*entity.Payment) presenter.I {
	return &listPresenter{
		payments: payments,
	}
}

// Exec method
func (p *listPresenter) Exec() ([]byte, error) {
	type listResponseItem struct {
		ID       uint32 `json:"id"`
		UserID   uint32 `json:"user_id"`
		Date     string `json:"date"`
		Price    uint32 `json:"price"`
		Category string `json:"category"`
		Memo     string `json:"memo"`
	}

	resp := make([]*listResponseItem, 0, len(p.payments))
	for _, payment := range p.payments {
		item := &listResponseItem{
			ID:       payment.ID,
			UserID:   payment.UserID,
			Date:     payment.Date.Format("2006/01/02"),
			Price:    payment.Price,
			Category: payment.Category,
			Memo:     payment.Memo,
		}
		resp = append(resp, item)
	}
	return json.Marshal(resp)
}
