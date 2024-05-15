package models

import (
	"encoding/json"

	"github.com/BrettS89/db"
	"github.com/BrettS89/util"
)

var modelName = "user"

type User struct {
	ID        string `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string `json:"firstName" bson:"firstName" validate:"required"`
	LastName  string `json:"lastName" bson:"lastName" validate:"required"`
	Email     string `json:"email" bson:"email" validate:"required"`
}

type PatchUserData struct {
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	Email     string `json:"email" bson:"email"`
}

func NewUser() User {
	return User{}
}

func Create(b []byte) (User, error) {
	util.ValidateData[User](b)

	var user User

	err := json.Unmarshal(b, &user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func Get(id string) {
	db.Client.Table(modelName).Get(id)
}

func Patch(id string, b []byte) {
	util.ValidateData[PatchUserData](b)
}

func Delete(id string) {

}
