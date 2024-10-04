package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Message struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	SenderID   primitive.ObjectID `bson:"sender_id"`
	ReceiverID primitive.ObjectID `bson:"receiver_id"`
	Content    string             `bson:"content"`
	Timestamp  time.Time          `bson:"timestamp"`
}
