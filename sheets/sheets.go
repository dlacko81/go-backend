package sheets

import (
	"context"
	"fmt"
	"log"

	"go-backend/handlers"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

const (
	spreadsheetId = "1MEyhm03JvbMPC4PTn7-NUraqYx6KZx0SH5xffCjbC2A"
	sheetName     = "Sheet1"
	credentials   = "credentials.json"
)

func getService() (*sheets.Service, context.Context, error) {
	ctx := context.Background()
	srv, err := sheets.NewService(ctx, option.WithCredentialsFile(credentials))
	if err != nil {
		log.Printf("Unable to retrieve Sheets client: %v", err)
		return nil, nil, err
	}
	return srv, ctx, nil
}

// Appends a new row
func AppendToSheet(data handlers.FormData) error {
	srv, ctx, err := getService()
	if err != nil {
		return err
	}

	row := []interface{}{
		data.ClientName,
		data.Date,
		data.Volume,
		data.Vintage,
		data.Technology,
		data.Country,
		data.Price,
		data.Comments,
	}

	valueRange := &sheets.ValueRange{
		Values: [][]interface{}{row},
	}

	_, err = srv.Spreadsheets.Values.Append(spreadsheetId, sheetName, valueRange).
		ValueInputOption("USER_ENTERED").
		InsertDataOption("INSERT_ROWS").
		Context(ctx).
		Do()

	if err != nil {
		log.Printf("Unable to append data: %v", err)
	}
	return err
}

// Reads all sheet data
func ReadSheet() ([][]interface{}, error) {
	srv, ctx, err := getService()
	if err != nil {
		return nil, err
	}

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, sheetName).Context(ctx).Do()
	if err != nil {
		log.Printf("Unable to retrieve data from sheet: %v", err)
		return nil, err
	}

	return resp.Values, nil
}

// Deletes a specific row (1-based index in Sheets, 0-based in code)
func DeleteRow(rowIndex int) error {
	srv, ctx, err := getService()
	if err != nil {
		return err
	}

	req := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			{
				DeleteDimension: &sheets.DeleteDimensionRequest{
					Range: &sheets.DimensionRange{
						SheetId:    0, // Default is 0 if only one sheet
						Dimension:  "ROWS",
						StartIndex: int64(rowIndex),     // 0-based inclusive
						EndIndex:   int64(rowIndex + 1), // exclusive
					},
				},
			},
		},
	}

	_, err = srv.Spreadsheets.BatchUpdate(spreadsheetId, req).Context(ctx).Do()
	if err != nil {
		log.Printf("Failed to delete row: %v", err)
	}
	return err
}

// Updates a specific row
func UpdateRow(rowIndex int, data handlers.FormData) error {
	srv, ctx, err := getService()
	if err != nil {
		return err
	}

	writeRange := fmt.Sprintf("%s!A%d:H%d", sheetName, rowIndex+1, rowIndex+1)

	row := []interface{}{
		data.ClientName,
		data.Date,
		data.Volume,
		data.Vintage,
		data.Technology,
		data.Country,
		data.Price,
		data.Comments,
	}

	valueRange := &sheets.ValueRange{
		Values: [][]interface{}{row},
	}

	_, err = srv.Spreadsheets.Values.Update(spreadsheetId, writeRange, valueRange).
		ValueInputOption("USER_ENTERED").
		Context(ctx).
		Do()

	if err != nil {
		log.Printf("Failed to update row: %v", err)
	}
	return err
}
