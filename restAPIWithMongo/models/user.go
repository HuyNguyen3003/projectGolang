package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID       bson.ObjectId
	username string
	password string
	age      int
}
