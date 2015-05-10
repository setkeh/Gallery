package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Image struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string
	Timestamp time.Time
	Data      string
}

type Person struct {
	ID int
}
