package models

import (
	"encoding/json"
	"io"

	"github.com/BrettS89/db"
	"github.com/BrettS89/util"
)

var modelName = "user"

type User struct {
	ID        string `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Email     string `json:"email,omitempty" bson:"email,omitempty"`
}

type CreateUserSchema struct {
	FirstName string `json:"firstName,omitempty" bson:"firstName,omitempty" validate:"required"`
	LastName  string `json:"lastName,omitempty" bson:"lastName,omitempty" validate:"required"`
	Email     string `json:"email,omitempty" bson:"email,omitempty" validate:"required"`
}

type PatchUserSchema struct {
	FirstName string `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Email     string `json:"email,omitempty" bson:"email,omitempty"`
}

func NewUser() User {
	return User{}
}

func (*User) Create(r io.ReadCloser) (User, error) {
	var user User

	b, err := io.ReadAll(r)

	if err != nil {
		return user, err
	}

	util.ValidateData[CreateUserSchema](b)

	err = json.Unmarshal(b, &user)

	if err != nil {
		return user, err
	}

	d, err := db.Client.Table(modelName).Create(user)

	if err != nil {
		return user, err
	}

	err = db.Client.Unmarshal(d, &user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (*User) Get(id string) (User, error) {
	var user User

	b, err := db.Client.Table(modelName).Get(id)

	if err != nil {
		return user, err
	}

	err = db.Client.Unmarshal(b, &user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (*User) Patch(id string, r io.ReadCloser) (User, error) {
	var user User

	b, err := io.ReadAll(r)

	if err != nil {
		return user, err
	}

	util.ValidateData[PatchUserSchema](b)

	err = json.Unmarshal(b, &user)

	if err != nil {
		return user, err
	}

	b, err = db.Client.Table(modelName).Patch(id, user)

	if err != nil {
		return user, err
	}

	db.Client.Unmarshal(b, &user)

	return user, nil
}

func Delete(id string) {

}
