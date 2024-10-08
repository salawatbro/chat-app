package repositories

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/salawatbro/chat-app/internal/interfaces"
	"github.com/salawatbro/chat-app/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type GroupRepository struct {
	db *mongo.Database
}

func NewGroupRepository(db *mongo.Database) interfaces.GroupRepositoryInterface {
	return &GroupRepository{
		db: db,
	}
}

func (repo *GroupRepository) FindAll(ctx *fiber.Ctx) ([]*models.Chat, error) {
	cursor, err := repo.db.Collection("chats").Find(ctx.Context(), bson.M{})
	if err != nil {
		return nil, err
	}
	var groups []*models.Chat
	if err := cursor.All(ctx.Context(), &groups); err != nil {
		return nil, err
	}
	return groups, nil
}

func (repo *GroupRepository) FindByID(ctx *fiber.Ctx, id string) (*models.Chat, error) {
	var group *models.Chat
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err // return an error if the conversion fails
	}
	pipeline := mongo.Pipeline{
		// Match stage to find the group by ID
		{{"$match", bson.M{"_id": objectID, "is_group": true}}},
		// Lookup stage to join with the user's collection
		{{"$lookup", bson.M{
			"from":         "users",      // The target collection
			"localField":   "created_by", // The field from the input documents
			"foreignField": "_id",        // The field from the documents of the "from" collection
			"as":           "owner",      // The name of the array fields to add to the output
		}}},
		// Unwind stage to deconstruct the owner array
		{{"$unwind", bson.M{"path": "$owner", "preserveNullAndEmptyArrays": true}}}, // Keep groups without an owner
	}
	cursor, err := repo.db.Collection("chats").Aggregate(ctx.Context(), pipeline)
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx.Context())
	if err != nil {
		return nil, err
	}
	if cursor.Next(ctx.Context()) {
		if err := cursor.Decode(&group); err != nil {
			return nil, err
		}
	} else {
		return nil, mongo.ErrNoDocuments // No group found
	}
	return group, nil
}

func (repo *GroupRepository) Create(ctx *fiber.Ctx, group *models.Chat) error {
	_, err := repo.db.Collection("chats").InsertOne(ctx.Context(), group)
	if err != nil {
		return err
	}
	return nil
}

func (repo *GroupRepository) Update(ctx *fiber.Ctx, id string, group *models.Chat) error {
	_, err := repo.db.Collection("chats").UpdateOne(ctx.Context(), bson.M{"_id": id}, group)
	if err != nil {
		return err
	}
	return nil
}

func (repo *GroupRepository) Delete(ctx *fiber.Ctx, id string) error {
	_, err := repo.db.Collection("chats").DeleteOne(ctx.Context(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
