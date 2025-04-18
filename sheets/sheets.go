package sheets

import (
	"context"
	"fmt"
	"log"

	"github.com/dlacko81/go-backend/models"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

const (
	spreadsheetId = "1MEyhm03JvbMPC4PTn7-NUraqYx6KZx0SH5xffCjbC2A"
	sheetName     = "Sheet1"
	credentials   = "credentials.json"
)

// getService creates and returns a Sheets API service client
func getService() (*sheets.Service, context.Context, error) {
	ctx := context.Background()
	srv, err := sheets.NewService(ctx, option.WithCredentialsFile(credentials))
	if err != nil {
		log.Printf("Unable to retrieve Sheets client: %v", err)
		return nil, nil, err
	}
	return srv, ctx, nil
}

// AppendToSheet appends a new row to the Google Sheets document
func AppendToSheet(data models.FormData) error {
	srv, ctx, err := getService()
	if err != nil {
		return err
	}

	// Prepare the row data
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

	// Prepare the ValueRange for the row
	valueRange := &sheets.ValueRange{
		Values: [][]interface{}{row},
	}

	// Append data to the sheet
	_, err = srv.Spreadsheets.Values.Append(spreadsheetId, sheetName, valueRange).
		ValueInputOption("USER_ENTERED").
		InsertDataOption("INSERT_ROWS").
		Context(ctx).
		Do()

	if err != nil {
		log.Printf("Unable to append data: %v", err)
		return err
	}
	return nil
}

// ReadSheet retrieves all data from the Google Sheets document
func ReadSheet() ([][]interface{}, error) {
	srv, ctx, err := getService()
	if err != nil {
		return nil, err
	}

	// Retrieve the sheet data
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, sheetName).Context(ctx).Do()
	if err != nil {
		log.Printf("Unable to retrieve data from sheet: %v", err)
		return nil, err
	}

	return resp.Values, nil
}

// DeleteRow deletes a specific row from the Google Sheets document
func DeleteRow(rowIndex int) error {
	srv, ctx, err := getService()
	if err != nil {
		return err
	}

	// Create the batch request for deleting the row
	req := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			{
				DeleteDimension: &sheets.DeleteDimensionRequest{
					Range: &sheets.DimensionRange{
						SheetId:    0,                      // Default sheet ID (usually 0)
						Dimension:  "ROWS",                 // Specify we're working with rows
						StartIndex: int64(rowIndex),       // 0-based inclusive index
						EndIndex:   int64(rowIndex + 1),   // exclusive end index (deletes one row)
					},
				},
			},
		},
	}

	// Execute the batch update (delete row)
	_, err = srv.Spreadsheets.BatchUpdate(spreadsheetId, req).Context(ctx).Do()
	if err != nil {
		log.Printf("Failed to delete row: %v", err)
		return err
	}
	return nil
}

// UpdateRow updates a specific row in the Google Sheets document
func UpdateRow(rowIndex int, data models.FormData) error {
	srv, ctx, err := getService()
	if err != nil {
		return err
	}

	// Specify the range to update (1-based index for Sheets)
	writeRange := fmt.Sprintf("%s!A%d:H%d", sheetName, rowIndex+1, rowIndex+1)

	// Prepare the row data to update
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

	// Prepare the ValueRange to send to Sheets API
	valueRange := &sheets.ValueRange{
		Values: [][]interface{}{row},
	}

	// Execute the update
	_, err = srv.Spreadsheets.Values.Update(spreadsheetId, writeRange, valueRange).
		ValueInputOption("USER_ENTERED").
		Context(ctx).
		Do()

	if err != nil {
		log.Printf("Failed to update row: %v", err)
		return err
	}
	return nil
}
