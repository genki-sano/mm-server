package gateway

import (
	"time"

	"github.com/genki-sano/mm-server/internal/entity"
)

// PaymentDataAccess type
type PaymentDataAccess interface {
	GetByDate(time.Time) ([]*entity.Payment, error)
}
