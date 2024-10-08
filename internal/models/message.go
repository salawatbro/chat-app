package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Attachment struct {
	FileName string `bson:"file_name"`
	FileURL  string `bson:"file_url"`
}

type Reaction struct {
	Like  primitive.ObjectID `bson:"like"`
	Heart primitive.ObjectID `bson:"heart"`
}

type Message struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	ChatID      primitive.ObjectID   `bson:"chat_id"`
	SenderID    primitive.ObjectID   `bson:"sender_id"`
	Text        string               `bson:"text"`
	Attachments []Attachment         `bson:"attachments"`
	SentAt      primitive.DateTime   `bson:"sent_at"`
	ReadBy      []primitive.ObjectID `bson:"read_by"`
	IsDeleted   bool                 `bson:"is_deleted"`
	Reactions   []Reaction           `bson:"reactions"`
}
