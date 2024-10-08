package utils

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserIdFromCtx(ctx *fiber.Ctx) primitive.ObjectID {
	return ctx.Locals("user_id").(primitive.ObjectID)
}
