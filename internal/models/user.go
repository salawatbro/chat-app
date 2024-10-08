package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name"`
	Email         string             `bson:"email"`
	Password      string             `bson:"password"`
	ProfilePicURL string             `bson:"profile_pic_url"`
	Status        string             `bson:"status"`
	CreatedAt     primitive.DateTime `bson:"created_at"`
}
