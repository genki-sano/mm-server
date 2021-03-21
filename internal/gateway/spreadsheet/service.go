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

func (s *service) getLastRow(ctx context.Context, spreadsheetID string, sheetTitle string) (int, error) {
	readRange := sheetTitle + "!A:A"
	valueRange, err := s.srv.Values.Get(spreadsheetID, readRange).Context(ctx).Do()
	if err != nil {
		return 0, err
	}
	return len(valueRange.Values), nil
}

func (s *service) isSheet(ctx context.Context, spreadsheetID string, sheetTitle string) (bool, error) {
	ss, err := s.srv.Get(spreadsheetID).Context(ctx).Do()
	if err != nil {
		return false, err
	}

	isSheet := false
	for _, sheet := range ss.Sheets {
		if sheet.Properties.Title == sheetTitle {
			isSheet = true
		}
	}
	return isSheet, nil
}

func (s *service) addSheet(ctx context.Context, spreadsheetID string, sheetIndex int64, sheetTitle string) error {
	req := sheets.Request{
		AddSheet: &sheets.AddSheetRequest{
			Properties: &sheets.SheetProperties{
				Index: sheetIndex,
				Title: sheetTitle,
			},
		},
	}

	rb := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{&req},
	}

	if _, err := s.srv.BatchUpdate(spreadsheetID, rb).Context(ctx).Do(); err != nil {
		return err
	}
	return nil
}
