package main

import (
	"file-upload-service/api"
	"github.com/gin-gonic/gin"
)

func main() {
	
	r := gin.Default()
	api.SetupRoutes(r)
	r.Run()

}