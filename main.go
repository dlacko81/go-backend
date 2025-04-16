package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/dlacko81/go-backend/sheets"
)

func main() {
	r := gin.Default()

	// GET: Fetch sheet data
	r.GET("/api/data", func(c *gin.Context) {
		data, err := sheets.GetSheetData()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, data)
	})

	// POST: Add new row to sheet
	r.POST("/api/data", func(c *gin.Context) {
		type FormInput struct {
			ClientName      string `json:"clientName"`
			TransactionDate string `json:"transactionDate"`
			Volume          string `json:"volume"`
			Vintage         string `json:"vintage"`
			Technology      string `json:"technology"`
			Country         string `json:"country"`
			Price           string `json:"price"`
			Comments        string `json:"comments"`
		}

		var input FormInput
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		row := []interface{}{
			input.ClientName,
			input.TransactionDate,
			input.Volume,
			input.Vintage,
			input.Technology,
			input.Country,
			input.Price,
			input.Comments,
		}

		if err := sheets.AppendRow(row); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to append to sheet"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	// Start the server on port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
