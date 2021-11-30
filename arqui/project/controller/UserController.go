package controller

import (
	u "arqui/project/model"
	rep "arqui/project/repository"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func PostUser(c *gin.Context) {
	var user u.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	id, err := rep.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func GetUsers(c *gin.Context) {
	var loadedUsers, err = rep.FindUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": loadedUsers})
}

func GetUserByNumber(c *gin.Context) {
	number := c.Param("number")
	var loadedUser, err = rep.FindUserByNumber(number)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ID": loadedUser.ID, "Name": loadedUser.Name})
}

func GetUsersWithNumber(c *gin.Context) {
	numbers := c.Param("numbers")
	var loadedUsers, err = rep.FindUsersWithNumber(numbers)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": loadedUsers})
}
