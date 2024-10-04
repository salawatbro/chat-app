package repositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/salawatbro/chat-app/internal/interfaces"
	"github.com/salawatbro/chat-app/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type AuthRepository struct {
	db *mongo.Database
}

func NewAuthRepository(db *mongo.Database) interfaces.AuthRepositoryInterface {
	return &AuthRepository{
		db: db,
	}
}

func (repo *AuthRepository) Register(ctx *fiber.Ctx, model models.User) (models.User, error) {
	model.ID = primitive.NewObjectID()
	user, err := repo.db.Collection("users").InsertOne(ctx.Context(), model)
	log.Println(user)
	if err != nil {
		return models.User{}, err
	}
	return model, nil
}

func (repo *AuthRepository) Login(ctx *fiber.Ctx, model models.User) (models.User, error) {
	//TODO implement me
	panic("implement me")
}
