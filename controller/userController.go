package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/inigoSutandyo/linkedin-copy-go/model"
)

func GetUserHandler(id uint) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		user := models.GetUserById(3)
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, gin.H{
			"message": "Some handler for my beautiful api",
			"user":    user,
		})
	}

	return gin.HandlerFunc(fn)
}
