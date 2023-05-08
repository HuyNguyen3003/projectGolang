package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	//"gopkg.in/mgo.v2"

	"github.com/HuyNguyen3003/WebServer/controllers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())

	//	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreatUser)
	//	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)

}

func getSession() *mongo.Client {
	s, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)

	}
	return s
}
