package interfaces

import (
	"github.com/gofiber/fiber/v2"
	"github.com/salawatbro/chat-app/internal/dto"
	"github.com/salawatbro/chat-app/internal/models"
)

type AuthRepositoryInterface interface {
	Register(ctx *fiber.Ctx, model models.User) (models.User, error)
	Login(ctx *fiber.Ctx, model models.Session) (models.Session, error)
	FindByEmail(ctx *fiber.Ctx, email string) (models.User, error)
	FindById(ctx *fiber.Ctx, id string) (models.User, error)
	FindSession(ctx *fiber.Ctx, token string) (models.Session, error)
}

type AuthServiceInterface interface {
	Register(ctx *fiber.Ctx, req *dto.RegisterDTO) error
	Login(ctx *fiber.Ctx, req *dto.LoginDTO) error
}

type AuthHandlerInterface interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
}
