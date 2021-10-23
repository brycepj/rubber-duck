package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./js-client/app/build/*.html")
	router.Static("/static", "./js-client/app/build/static")
	router.GET("/", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "index.html", nil)
	})
	router.Run("localhost:8080")
}
