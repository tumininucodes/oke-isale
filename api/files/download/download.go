package download

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Download(c *gin.Context) {
	filename := c.Param("filename")

		filePath := "storage/" + filename

		_, err := os.Stat(filePath)
		if os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
			return
		}

		file, err := os.Open(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
			return
		}
		defer file.Close()

		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Disposition", "attachment; filename="+filename)
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Expires", "0")
		c.Header("Cache-Control", "must-revalidate")
		c.Header("Pragma", "public")

		io.Copy(c.Writer, file)
}