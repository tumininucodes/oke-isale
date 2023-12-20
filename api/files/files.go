package files

import "github.com/gin-gonic/gin"

func FetchFiles(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get all users"})
}

func FetchFile(c *gin.Context) {

}