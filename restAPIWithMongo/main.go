package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"

	"github.com/HuyNguyen3003/WebServer/controllers"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.getUser)
	r.POST("/user", uc.creatUser)
	r.DELETE("/user/:id", uc.deleteUser)
	http.ListenAndServe("localhost:9000", r)

}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)

	}
	return s
}
