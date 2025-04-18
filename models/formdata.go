package models

type FormData struct {
	ClientName string  `json:"clientName"`
	Date       string  `json:"date"`
	Volume     float64 `json:"volume"`
	Vintage    string  `json:"vintage"`
	Technology string  `json:"technology"`
	Country    string  `json:"country"`
	Price      float64 `json:"price"`
	Comments   string  `json:"comments"`
	RowIndex   int     `json:"rowIndex"` // Used for updates/deletes
}
