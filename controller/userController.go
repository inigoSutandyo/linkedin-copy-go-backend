package controller

import (
	"fmt"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/inigoSutandyo/linkedin-copy-go/model"
)

func GetUserByIdHandler(c *gin.Context) {
	id := new(big.Int)
	_, err := fmt.Sscan(c.Param("id"), id)
	if err != nil {
		fmt.Println("error scanning value:", err)
	} else {
		fmt.Println(id)
	}

	user := models.GetUserById(*id)
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "Some handler for my beautiful api",
		"user":    user,
	})
}

func GetAllUsersHandler(c *gin.Context) {
	users := models.GetAllUsers()
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"users":   users,
	})
}
