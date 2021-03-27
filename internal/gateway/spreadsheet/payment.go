package spreadsheet

import (
	"context"
	"log"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/genki-sano/mm-server/internal/entity"
	"github.com/genki-sano/mm-server/internal/gateway"
	"google.golang.org/api/sheets/v4"
)

const location = "Asia/Tokyo"

type paymentRepository struct {
	srv           *service
	spreadsheetID string
	ctx           context.Context
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
		ctx:           ctx,
	}
}

func (r *paymentRepository) GetByDate(t time.Time) ([]*entity.Payment, error) {
	readRange := t.Format("2006-01") + "!A:J"
	valueRange, err := r.srv.get(r.ctx, r.spreadsheetID, readRange)
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
		userType, err := strconv.Atoi(items[1].(string))
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
			UserType: uint8(userType),
			Category: items[2].(string),
			Price:    uint32(price),
			Date:     date,
			Memo:     items[5].(string),
		}
		payments = append(payments, payment)
	}
	return payments, nil
}

func (r *paymentRepository) Insert(payment *entity.Payment) error {
	if err := r.isFirstInTheMonth(payment.Date); err != nil {
		return err
	}

	lastRow, err := r.srv.getLastRow(r.ctx, r.spreadsheetID, "payments")
	if err != nil {
		return err
	}

	uid := lastRow
	row := lastRow + 1

	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc

	now := time.Now()

	item := make([]interface{}, 0, 11)
	item = append(item, uid)
	item = append(item, payment.UserType)
	item = append(item, payment.Category)
	item = append(item, payment.Price)
	item = append(item, payment.Date.Format("2006/01/02"))
	item = append(item, payment.Memo)
	item = append(item, payment.UserType)
	item = append(item, payment.UserType)
	item = append(item, now.Format("2006/01/02 15:04:05"))
	item = append(item, now.Format("2006/01/02 15:04:05"))
	item = append(item, "=DATEVALUE(E"+strconv.Itoa(row)+")")

	values := make([][]interface{}, 0, 1)
	values = append(values, item)

	appendRange := "payments!A:J"
	rb := &sheets.ValueRange{
		Values: values,
	}

	if _, err := r.srv.insert(r.ctx, r.spreadsheetID, appendRange, rb); err != nil {
		return err
	}

	return nil
}

func (r *paymentRepository) isFirstInTheMonth(date time.Time) error {
	sheetTitle := date.Format("2006-01")

	isSheet, err := r.srv.isSheet(r.ctx, r.spreadsheetID, sheetTitle)
	if err != nil {
		return err
	}

	if isSheet {
		return nil
	}

	if err := r.srv.addSheet(r.ctx, r.spreadsheetID, 2, sheetTitle); err != nil {
		return err
	}

	thisT := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
	nextT := thisT.AddDate(0, 1, 0)

	thisMonth := strconv.FormatFloat(math.Floor(timeToSerialValue(thisT, false)), 'f', -1, 64)
	nextMonth := strconv.FormatFloat(math.Floor(timeToSerialValue(nextT, false)), 'f', -1, 64)

	item := make([]interface{}, 0, 1)
	item = append(item, "=QUERY(payments!A2:K,\"select A,B,C,D,E,F,G,H,I,J where K >= "+thisMonth+" and K < "+nextMonth+" order by K desc\")")

	values := make([][]interface{}, 0, 1)
	values = append(values, item)

	appendRange := sheetTitle + "!A1"
	rb := &sheets.ValueRange{
		Values: values,
	}

	if _, err := r.srv.insert(r.ctx, r.spreadsheetID, appendRange, rb); err != nil {
		return err
	}
	return nil
}
