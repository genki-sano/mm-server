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

// @see https://developers.google.com/sheets/api/reference/rest/v4/spreadsheets.values/append
func (s *service) insert(ctx context.Context, spreadsheetID string, appendRange string, valueRange *sheets.ValueRange) (*sheets.AppendValuesResponse, error) {
	return s.srv.Values.Append(spreadsheetID, appendRange, valueRange).InsertDataOption("INSERT_ROWS").ValueInputOption("USER_ENTERED").Context(ctx).Do()
}

// @see https://developers.google.com/sheets/api/reference/rest/v4/spreadsheets.values/update
func (s *service) update(ctx context.Context, spreadsheetID string, updateRange string, valueRange *sheets.ValueRange) (*sheets.UpdateValuesResponse, error) {
	return s.srv.Values.Update(spreadsheetID, updateRange, valueRange).ValueInputOption("USER_ENTERED").Context(ctx).Do()
}
