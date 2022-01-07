package app

import "github.com/gin-gonic/gin"

var router = gin.Default()

func Run() {
	mapUrls()

	router.Run(":8000")
}
