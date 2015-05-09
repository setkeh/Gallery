package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"os"
)

var (
	mdbSession *mgo.Session
)

type Image struct {
	Name string
}

type Person struct {
	Id int
}

type Config struct {
	Mongo string
}

func db(d string, config *Config) {
	// Save the connection session so that we can access the DB as a global
	mdbSession, err := mgo.Dial(config.Mongo)
	if err != nil {
		panic(err)
	}
	defer mdbSession.Close()

	c := mdbSession.DB("Gallery").C("images")
	image := &Image{d}
	err = c.Insert(image.Name)
	if err != nil {
		log.Fatal(err)
	}

	result := Person{1}
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
