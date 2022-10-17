package controllers

import (
	"fmt"
	"net/http"

	"RefugeWallet/database"

	"github.com/gin-gonic/gin"
)

type userRegister struct {
	Login string `json:"Login" binding:"required"`
	Pass  string `json:"Pass" binding:"required"`
}

func Register(c *gin.Context) {
	var user userRegister
	var us database.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len([]rune(user.Pass)) < 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password should contain 6 characters and more"})
		return
	}

	fmt.Printf("Struct: %v\n", user)
	us.Login = user.Login
	us.Pass = user.Pass

	_, err := database.AddUser(us)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Registration Success!"})
}
