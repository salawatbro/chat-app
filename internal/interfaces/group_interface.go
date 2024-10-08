package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/salawatbro/chat-app/internal/dto"
	"github.com/salawatbro/chat-app/internal/models"
)

type GroupRepositoryInterface interface {
	FindAll(ctx *fiber.Ctx) ([]*models.Chat, error)
	FindByID(ctx *fiber.Ctx, id string) (*models.Chat, error)
	Create(ctx *fiber.Ctx, group *models.Chat) error
	Update(ctx *fiber.Ctx, id string, group *models.Chat) error
	Delete(ctx *fiber.Ctx, id string) error
}

type GroupServiceInterface interface {
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx, id string) error
	Create(ctx *fiber.Ctx, req *dto.GroupDTO) error
	Update(ctx *fiber.Ctx, id string, req *dto.GroupDTO) error
	Delete(ctx *fiber.Ctx, id string) error
}

type GroupHandlerInterface interface {
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
