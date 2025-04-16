package sheets

import (
	"context"
	"log"
	"google.golang.org/api/sheets/v4"
	"google.golang.org/api/option"
)

// Function to get data from Google Sheets
func GetSheetData() ([]interface{}, error) {
	ctx := context.Background()

	// Create a new Sheets API client
	srv, err := sheets.NewService(ctx, option.WithAPIKey("YOUR_GOOGLE_API_KEY"))
	if err != nil {
		log.Fatalf("Unable to create Sheets service: %v", err)
		return nil, err
	}

	// Specify the spreadsheet ID and the range of cells
	spreadsheetId := "https://docs.google.com/spreadsheets/d/1MEyhm03JvbMPC4PTn7-NUraqYx6KZx0SH5xffCjbC2A/edit?usp=sharing" // Replace with your actual spreadsheet ID
	readRange := "Sheet1!A1:F10"           // Adjust the range as needed

	// Retrieve data from Google Sheets
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
		return nil, err
	}

	return resp.Values, nil
}
