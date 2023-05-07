package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/HuyNguyen3003/WebServer/models"
)

type UserController struct {
	Session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}

}

func (uc UserController) getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)
	u := models.User{}
	if err := uc.Session.DB("curdGolang").C("user").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}
	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "%s\n", uj)
}

func (uc UserController) createUser(w http.ResponseWriter, r *http.Request) {

	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)
	u.ID = bson.NewObjectId()

	uc.Session.DB("curdGolang").C("user").Insert(u)

	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "%s\n", uj)

}

func (uc UserController) deleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)
	u := models.User{}
	if err := uc.Session.DB("curdGolang").C("user").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}
	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "%s\n", uj)
}
