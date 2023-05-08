package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"

	"github.com/HuyNguyen3003/WebServer/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	Session *mongo.Client
}

func NewUserController(s *mongo.Client) *UserController {
	return &UserController{s}

}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)
	u := models.User{}
	if err := uc.Session.Database("curdGolang").Collection("users").FindOne(context.TODO(), bson.M{"_id": oid}); err != nil {
		w.WriteHeader(404)
		return
	}
	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "\n", uj)
}

func (uc UserController) CreatUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	u := models.User{}

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u.ID = bson.NewObjectId()

	uc.Session.Database("curdGolang").Collection("users").InsertOne(context.TODO(), u)

	uj, err := json.Marshal(u)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "\n", string(uj))

}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)
	u := models.User{}
	deleteResult, err := uc.Session.Database("curdGolang").Collection("users").DeleteOne(context.TODO(), bson.M{"_id": oid})
	if err != nil {

		w.WriteHeader(404)
		return
	}
	deletedCount := deleteResult.DeletedCount
	fmt.Printf("Đã xóa %d bản ghi", deletedCount)

	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "\n", uj)
}
