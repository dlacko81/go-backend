package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/dlacko81/go-backend/sheets"
)

func main() {
	r := gin.Default()

	// Allow your frontend domain through CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://go-frontend-mu.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	// GET: Fetch sheet data
	r.GET("/api/data", func(c *gin.Context) {
		data, err := sheets.GetSheetData()
		if err != nil {
			log.Printf("GetSheetData failed: %v", err)
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
			ClientDirection string `json:"clientDirection"`
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
			input.ClientDirection,
			input.Volume,
			input.Vintage,
			input.Technology,
			input.Country,
			input.Price,
			input.Comments,
		}

		if err := sheets.AppendRow(row); err != nil {
			log.Printf("AppendRow failed: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	// PUT: Update an existing row
	r.PUT("/api/data", func(c *gin.Context) {
		type UpdateInput struct {
			RowIndex        int    `json:"rowIndex"`
			ClientName      string `json:"clientName"`
			TransactionDate string `json:"transactionDate"`
			ClientDirection string `json:"clientDirection"`
			Volume          string `json:"volume"`
			Vintage         string `json:"vintage"`
			Technology      string `json:"technology"`
			Country         string `json:"country"`
			Price           string `json:"price"`
			Comments        string `json:"comments"`
		}

		var input UpdateInput
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		row := []interface{}{
			input.ClientName,
			input.TransactionDate,
			input.ClientDirection,
			input.Volume,
			input.Vintage,
			input.Technology,
			input.Country,
			input.Price,
			input.Comments,
		}

		if err := sheets.UpdateRow(input.RowIndex, row); err != nil {
			log.Printf("UpdateRow failed: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "updated"})
	})

	// DELETE: Delete a row
	r.DELETE("/api/data/:row", func(c *gin.Context) {
		rowStr := c.Param("row")
		rowIndex, err := strconv.Atoi(rowStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid row index"})
			return
		}

		if err := sheets.DeleteRow(rowIndex); err != nil {
			log.Printf("DeleteRow failed: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "deleted"})
	})

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
