package spreadsheet

import (
	"context"

	"google.golang.org/api/sheets/v4"
)

type service struct {
	srv *sheets.SpreadsheetsService
}

func newService(ctx context.Context) (*service, error) {
	// Spreadsheetの書き込み権限まで付与されている想定
	srv, err := sheets.NewService(ctx)
	if err != nil {
		return nil, err
	}
	s := &service{
		srv: srv.Spreadsheets,
	}
	return s, nil
}

// @see https://developers.google.com/sheets/api/reference/rest/v4/spreadsheets.values/get
func (s *service) get(ctx context.Context, spreadsheetID string, readRange string) (*sheets.ValueRange, error) {
	return s.srv.Values.Get(spreadsheetID, readRange).Context(ctx).Do()
}
