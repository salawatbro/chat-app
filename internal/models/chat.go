package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Participant struct {
	UserId   primitive.ObjectID `bson:"user_id"`
	JoinedAt primitive.DateTime `bson:"joined_at"`
}

type LastMessage struct {
	MessageID primitive.ObjectID `bson:"message_id"`
	SentAt    primitive.DateTime `bson:"sent_at"`
	Text      string             `bson:"text"`
}

type Chat struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	IsGroup      bool               `bson:"is_group"`
	IsPublic     bool               `bson:"is_public"`
	ChatName     string             `bson:"chat_name"`
	Participants []Participant      `bson:"participants"`
	LastMessage  *LastMessage       `bson:"last_message"`
	CreatedBy    primitive.ObjectID `bson:"created_by"`
	CreatedAt    primitive.DateTime `bson:"created_at"`
	UpdatedAt    primitive.DateTime `bson:"updated_at"`

	Owner User `bson:"-"`
}
