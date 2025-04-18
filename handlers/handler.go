package handlers

import (
	"net/http"
	"strconv"

	"github.com/dlacko81/go-backend/sheets"  // Ensure that the import path matches your module structure
	"github.com/gin-gonic/gin"
)

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

// SubmitData handles submitting new data and appending it to the sheet
func SubmitData(c *gin.Context) {
	var data FormData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := sheets.AppendToSheet(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write to sheet"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// GetData retrieves data from the sheet and returns it
func GetData(c *gin.Context) {
	records, err := sheets.ReadSheet()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read data"})
		return
	}
	c.JSON(http.StatusOK, records)
}

// DeleteData deletes a specific row in the sheet based on the row index
func DeleteData(c *gin.Context) {
	rowStr := c.Param("row")
	rowIndex, err := strconv.Atoi(rowStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid row index"})
		return
	}
	err = sheets.DeleteRow(rowIndex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete row"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

// UpdateData updates a specific row in the sheet based on the row index
func UpdateData(c *gin.Context) {
	var data FormData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := sheets.UpdateRow(data.RowIndex, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update row"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}
