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
		c.JSON(http.StatusOK, gin.H{"data": data})
	})

	// POST: Add new row to sheet
	r.POST("/api/data", func(c *gin.Context) {
		var row []interface{}
		if err := c.BindJSON(&row); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}

		if err := sheets.AppendRow(row); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to append to sheet"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
