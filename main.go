package main

import (
	"golang_cicd/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	router.GET("/", controllers.ChineseMoneyIndex)
	router.Run(":8081")
}
