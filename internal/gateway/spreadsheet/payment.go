package spreadsheet

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/genki-sano/mm-server/internal/entity"
	"github.com/genki-sano/mm-server/internal/gateway"
)

type paymentRepository struct {
	srv           *service
	spreadsheetID string
}

// NewPaymentRepository method
func NewPaymentRepository() gateway.PaymentDataAccess {
	ctx := context.Background()
	srv, err := newService(ctx)
	if err != nil {
		// TODO: エラーハンドリングを悩み中
		log.Fatalf(err.Error())
	}
	id := os.Getenv("GOOGLE_SPREDSHEET_ID")
	return &paymentRepository{
		srv:           srv,
		spreadsheetID: id,
	}
}

func (r *paymentRepository) GetByDate(t time.Time) ([]*entity.Payment, error) {
	readRange := t.Format("2006-01") + "!A:J"
	valueRange, err := r.srv.get(r.spreadsheetID, readRange)
	if err != nil {
		return nil, err
	}

	payments := make([]*entity.Payment, 0)
	if len(valueRange.Values) == 0 {
		return payments, nil
	}

	for _, items := range valueRange.Values {
		id, err := strconv.Atoi(items[0].(string))
		if err != nil {
			// TODO: エラーハンドリングを悩み中
			log.Fatalf(err.Error())
		}
		userID, err := strconv.Atoi(items[1].(string))
		if err != nil {
			// TODO: エラーハンドリングを悩み中
			log.Fatalf(err.Error())
		}
		price, err := strconv.Atoi(items[3].(string))
		if err != nil {
			// TODO: エラーハンドリングを悩み中
			log.Fatalf(err.Error())
		}
		date, err := time.ParseInLocation("2006/01/02", items[4].(string), time.Local)
		if err != nil {
			log.Fatalf(err.Error())
		}
		payment := &entity.Payment{
			ID:       uint32(id),
			UserID:   uint32(userID),
			Category: items[2].(string),
			Price:    uint32(price),
			Date:     date,
			Memo:     items[5].(string),
		}
		payments = append(payments, payment)
	}
	return payments, nil
}
