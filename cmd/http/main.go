package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.POST("/content", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello go!",
		})
	})

	r.Run()
}
