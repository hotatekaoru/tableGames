package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func A01G01(c *gin.Context) {

	c.HTML(http.StatusOK, "a01.html", gin.H{})

}
