package presenter

import "github.com/genki-sano/mm-server/internal/entity"

// PaymentListFactory type
type PaymentListFactory interface {
	New([]*entity.Payment) I
}
