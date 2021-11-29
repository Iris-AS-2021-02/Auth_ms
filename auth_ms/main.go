package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.POST("/user", postUser)
	router.GET("/user/number/:number", getUserByNumber)
	router.GET("/user", getUsers)
	router.Run("localhost:8080")
}
