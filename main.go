package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
	"time"
)

var (
	mdbSession *mgo.Session
)

type Config struct {
	Mongo      string
	Database   string
	Collection string
}

func db(d string, config *Config) {
	// Save the connection session so that we can access the DB as a global
	mdbSession, err := mgo.Dial(config.Mongo)
	if err != nil {
		panic(err)
	}
	defer mdbSession.Close()

	c := mdbSession.DB(config.Database).C(config.Collection)
	imgObjId := bson.NewObjectId()
	imgData := base64.StdEncoding.EncodeToString([]byte("image data converted to base64 string"))
	image := &Image{imgObjId, "Test", time.Now(), imgData}
	err = c.Insert(image)
	if err != nil {
		log.Fatal(err)
	}

	result := Person{1}
	err = c.Find(bson.M{"image": d}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

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
	server.GET("/stats", GetAllStats)
	server.GET("/image", GetAllImages)
	server.POST("/image", CreateImage)
	server.GET("/image/:id", GetImage)
	server.POST("/image/:id", UpdateImage)
	server.GET("/", PingHandler) // Change the handler to taste

	// Start server on localhost:3000
	server.Run(":3000")
}
