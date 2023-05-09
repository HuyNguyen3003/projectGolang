package initializer

import "github.com/kamva/mgm/v3"

type User struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model.
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Password         int    `json:"password" bson:"password"`
}
