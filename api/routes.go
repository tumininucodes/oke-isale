package api

import (
	"file-upload-service/api/files"
	"file-upload-service/api/files/download"
	"file-upload-service/api/files/upload"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	filesRoute := r.Group("api/files")

	filesRoute.GET("/", files.FetchFiles)
	filesRoute.GET("/download/:filename", download.Download)
	filesRoute.POST("/upload", upload.Upload)

}