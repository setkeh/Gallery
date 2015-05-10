package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "PONG!"})
}

func GetAllStats(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "statistics go here"})
}

func GetAllImages(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "all images go here"})
}

func GetImage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "one image goes here"})
}

func CreateImage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "ID of the image we created goes here"})
}

func UpdateImage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "ID of the image we updated goes here"})
}
