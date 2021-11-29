package main

import (
	c "arqui/project/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.POST("/user", c.PostUser)
	router.GET("/user/number/:number", c.GetUserByNumber)
	router.GET("/user", c.GetUsers)
	router.Run("localhost:8080")
}
