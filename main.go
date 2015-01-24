package main

import (
	//"encoding/base64"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//"net/http"
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

func main() {
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
