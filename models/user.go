package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"id" `
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
}

func (user *User) Login() bool {
	err := db.User.FindOne(context.TODO(), user).Decode(user)
	return err == nil
}
