package main

import (
	"fmt"

	"RefugeWallet/controllers"
	"RefugeWallet/database"
	"RefugeWallet/middleware"

	"github.com/gin-gonic/gin"
)

/*
func authPage(c *gin.Context) {
	login := c.Query("Login")
	pass := c.Query("Pass")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	fmt.Printf("HI it's header \n%s\n%s\n", login, pass)
	fmt.Printf("Client IP: %s\n", c.ClientIP())
	c.Data(http.StatusOK, "text/html", []byte("Login endpoint!\n"))
}
*/

func main() {

	database.Connect()
	router := gin.Default()

	public := router.Group("/api")

	public.POST("/registration", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := router.Group("/api/wallet")
	protected.Use(middleware.JwtAuth())
	//	protected.GET("/", controllers.WalletStatus)
	protected.GET("/user", controllers.CurrentUser)

	fmt.Println("Started!")
	router.Run("localhost:8080")
}
