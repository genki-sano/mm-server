package spreadsheet

import (
	"context"

	"google.golang.org/api/sheets/v4"
)

type service struct {
	s *sheets.SpreadsheetsService
}

func newService(ctx context.Context) (*service, error) {
	// Spreadsheetの書き込み権限まで付与されている想定
	srv, err := sheets.NewService(ctx)
	if err != nil {
		return nil, err
	}
	c := &service{
		s: srv.Spreadsheets,
	}
	return c, nil
}

func (s *service) get(spreadsheetID string, readRange string) (*sheets.ValueRange, error) {
	return s.s.Values.Get(spreadsheetID, readRange).Do()
}
