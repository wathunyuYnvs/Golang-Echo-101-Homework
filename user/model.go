package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Name     string             `json:"name"`
	Age      int                `json:"age"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}

type Users []User
