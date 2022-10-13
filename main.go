package main

import (
	"fmt"
	//	"log"
	//	"ner/http"
	"database/sql"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func main() {
	router := gin.Default()
	fmt.Println("Started!")
	router.Run("localhost:8080")
}
