package sheets

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

const spreadsheetId = "1MEyhm03JvbMPC4PTn7-NUraqYx6KZx0SH5xffCjbC2A"
const sheetRange = "Sheet1"

type FormData struct {
	ClientName string
	Date       string
	Volume     float64
	Vintage    string
	Technology string
	Country    string
	Price      float64
	Comments   string
}

func getService() (*sheets.Service, error) {
	ctx := context.Background()
	b, err := ioutil.ReadFile("credentials.json") // Downloaded from Google Cloud
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	config := option.WithCredentialsJSON(b)
	srv, err := sheets.NewService(ctx, config)
	return srv, err
}

func AppendToSheet(data FormData) error {
	srv, err := getService()
	if err != nil {
		return err
	}

	row := []interface{}{data.ClientName, data.Date, data.Volume, data.Vintage, data.Technology, data.Country, data.Price, data.Comments}
	_, err = srv.Spreadsheets.Values.Append(spreadsheetId, sheetRange, &sheets.ValueRange{
		Values: [][]interface{}{row},
	}).ValueInputOption("USER_ENTERED").Do()

	return err
}

func ReadSheet() ([][]interface{}, error) {
	srv, err := getService()
	if err != nil {
		return nil, err
	}
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, sheetRange).Do()
	if err != nil {
		return nil, err
	}
	return resp.Values, nil
}
