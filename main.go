package main

import (
	//"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"os"
	//"time"
)

type Image struct {
	image string
}

type Config struct {
	Mongo string
}

func db(d string) {
	session, err := mgo.Dial(config.Mongo)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("Gallery").C("images")
	err = c.Insert(&image{d})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"image": d}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

}

// Request handler sets the response context
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "PONG!"})
}

func init() {
	file, err := os.Open("config.json")

	if err != nil {
		fmt.Println("Couldn't read config file, dying...")
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}
	decoder.Decode(&config)
}

func main() {
	// REST server
	server := gin.Default()

	// Routes
	server.GET("/ping", PingHandler)
	server.GET("/", PingHandler) // Change the handler to taste

	// Start server on localhost:3000
	server.Run(":3000")
}
