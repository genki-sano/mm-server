package valueobject

import (
	"time"
)

// PaymentDate type
type PaymentDate struct {
	value time.Time
}

// NewPaymentDate method
func NewPaymentDate(value string) (*PaymentDate, error) {
	if value == "" {
		return nil, newRequiredError("日付")
	}

	date, err := time.ParseInLocation("2006-01", value, time.Local)
	if err != nil {
		return nil, newFormatError("日付")
	}

	return &PaymentDate{value: date}, nil
}

// Get method
func (vo *PaymentDate) Get() time.Time {
	return vo.value
}
