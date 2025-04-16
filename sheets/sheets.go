package sheets

import (
	"context"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// Reads data from the Google Sheet
func GetSheetData() ([][]interface{}, error) {
	ctx := context.Background()

	// Assumes "credentials.json" is in the root of your Render project directory
	srv, err := sheets.NewService(ctx, option.WithCredentialsFile("credentials.json"))
	if err != nil {
		log.Printf("Unable to retrieve Sheets client: %v", err)
		return nil, err
	}

	spreadsheetId := "1MEyhm03JvbMPC4PTn7-NUraqYx6KZx0SH5xffCjbC2A"
	readRange := "Sheet1" // Change this if your sheet has a different name

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Printf("Unable to retrieve data from sheet: %v", err)
		return nil, err
	}

	return resp.Values, nil
}

// Appends a new row to the Google Sheet
func AppendRow(row []interface{}) error {
	ctx := context.Background()

	// Load the Sheets service
	srv, err := sheets.NewService(ctx, option.WithCredentialsFile("credentials.json"))
	if err != nil {
		log.Printf("Unable to retrieve Sheets client: %v", err)
		return err
	}

	spreadsheetId := "1MEyhm03JvbMPC4PTn7-NUraqYx6KZx0SH5xffCjbC2A"
	writeRange := "Sheet1" // Adjust if needed

	valueRange := &sheets.ValueRange{
		Values: [][]interface{}{row},
	}

	_, err = srv.Spreadsheets.Values.Append(spreadsheetId, writeRange, valueRange).
		ValueInputOption("USER_ENTERED").
		InsertDataOption("INSERT_ROWS").
		Context(ctx).
		Do()

	if err != nil {
		log.Printf("Unable to append data to sheet: %v", err)
		return err
	}

	return nil
}
