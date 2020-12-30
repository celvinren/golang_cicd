package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChineseMoneyIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "chinese_money_index.html", gin.H{})
}
