package upload

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if file.Size > (8 << 20) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds the limit (8MB)"})
		return
	}

	err = c.SaveUploadedFile(file, "storage/"+file.Filename)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}