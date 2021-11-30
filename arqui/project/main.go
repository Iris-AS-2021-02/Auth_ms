package main

import (
	c "arqui/project/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/user", c.PostUser)
	router.GET("/user/:number", c.GetUserByNumber)
	router.GET("/user", c.GetUsers)
	router.GET("/user/find/:numbers", c.GetUsersWithNumber)
	router.Run("0.0.0.0:8080")
}
