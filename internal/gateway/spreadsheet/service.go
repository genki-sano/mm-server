package spreadsheet

import (
	"context"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type service struct {
	s *sheets.SpreadsheetsService
}

func newService(ctx context.Context) (*service, error) {
	srv, err := sheets.NewService(ctx, option.WithScopes(sheets.SpreadsheetsReadonlyScope))
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
