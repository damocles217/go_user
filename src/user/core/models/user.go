package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name      string             `bson:"name,omitempty" json:"name,omitempty"`
	Lastname  string             `bson:"lastname,omitempty" json:"lastname,omitempty"`
	Email     string             `bson:"email,omitempty" json:"email,omitempty"`
	Password  string             `bson:"password,omitempty" json:"password,omitempty"`
	UserID    string             `bson:"userId,omitempty" json:"userId,omitempty"`
	CodeAuth  string             `bson:"code_auth,omitempty" json:"code_auth,omitempty"`
	Admin     int                `bson:"admin,omitempty" json:"admin,omitempty"`
	Gender    int                `bson:"gender,omitempty" json:"gender,omitempty"`
	UrlPhoto  string             `bson:"url_photo,omitempty" json:"url_photo,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

type UserCreate struct {
	Email           string `bson:"email,omitempty" json:"email,omitempty"`
	Name            string `bson:"name,omitempty" json:"name,omitempty"`
	Lastname        string `bson:"lastname,omitempty" json:"lastname,omitempty"`
	Password        string `bson:"password,omitempty" json:"password,omitempty"`
	ConfirmPassword string `json:"confirm_password,omitempty"`
	Gender          int    `bson:"gender,omitempty" json:"gender,omitempty"`
}
