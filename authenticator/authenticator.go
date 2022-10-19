package controllers

import (
	"fmt"
	"net/http"

	"RefugeWallet/database"
	"RefugeWallet/token"

	"github.com/gin-gonic/gin"
)

type userRegister struct {
	U_Id  uint
	Login string `json:"Login" binding:"required"`
	Pass  string `json:"Pass" binding:"required"`
}

func Login(c *gin.Context) {
	var login userRegister
	var us database.User

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	us.U_Id = login.U_Id
	us.Login = login.Login
	us.Pass = login.Pass

	token, err := database.LoginCheck(us)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

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

	//fmt.Printf("Struct: %v\n", user)
	us.Login = user.Login
	us.Pass = user.Pass

	_, err := database.AddUser(us)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Registration Success!", "msg": (fmt.Sprintf("Welcome %s!!!", us.Login))})
}

func CurrentUser(c *gin.Context) {
	userID, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := database.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}
