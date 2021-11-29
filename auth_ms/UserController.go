package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func postUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	id, err := createUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func getUsers(c *gin.Context) {
	var loadedUsers, err = findUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": loadedUsers})
}

func getUserByNumber(c *gin.Context) {
	number := c.Param("number")
	var loadedUser, err = findUserByNumber(number)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ID": loadedUser.ID, "Name": loadedUser.Name})
}
