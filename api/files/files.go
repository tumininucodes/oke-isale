package files

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func FetchFiles(c *gin.Context) {
	
	uploadsDir := "storage/"

	files, err := os.ReadDir(uploadsDir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read directory"})
		return
	}

	var fileList []gin.H

	for _, file := range files {
		fileInfo, err := file.Info()
		if err != nil {
			fmt.Println("Error getting file info:", err)
			continue
		}

		fileData := gin.H{
			"name": file.Name(),
			"size": fmt.Sprintf("%.2fKB", float64(fileInfo.Size())/1024), // Convert size to kilobytes
		}

		fileList = append(fileList, fileData)
	}

	c.JSON(http.StatusOK, fileList)
}