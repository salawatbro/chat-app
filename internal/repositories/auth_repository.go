package repositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/salawatbro/chat-app/internal/interfaces"
	"github.com/salawatbro/chat-app/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	_, err := repo.db.Collection("users").InsertOne(ctx.Context(), model)
	if err != nil {
		return models.User{}, err
	}
	user, err := repo.FindById(ctx, model.ID.Hex())
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (repo *AuthRepository) FindByEmail(ctx *fiber.Ctx, email string) (models.User, error) {
	var user models.User
	err := repo.db.Collection("users").FindOne(ctx.Context(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (repo *AuthRepository) FindById(ctx *fiber.Ctx, id string) (models.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, err
	}
	var user models.User
	err = repo.db.Collection("users").FindOne(ctx.Context(), bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (repo *AuthRepository) Login(ctx *fiber.Ctx, model models.Session) (models.Session, error) {
	_, err := repo.db.Collection("sessions").InsertOne(ctx.Context(), model)
	if err != nil {
		return models.Session{}, err
	}
	session, err := repo.FindSession(ctx, model.Token)
	if err != nil {
		return models.Session{}, err
	}
	return session, nil
}

func (repo *AuthRepository) FindSession(ctx *fiber.Ctx, token string) (models.Session, error) {
	var session models.Session
	err := repo.db.Collection("sessions").FindOne(ctx.Context(), bson.M{"token": token}).Decode(&session)
	if err != nil {
		return models.Session{}, err
	}
	return session, nil
}
