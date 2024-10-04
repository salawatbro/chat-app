package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/salawatbro/chat-app/internal/dto"
	"github.com/salawatbro/chat-app/internal/interfaces"
	"github.com/salawatbro/chat-app/pkg/utils"
)

type AuthHandler struct {
	service interfaces.AuthServiceInterface
}

func NewAuthHandler(service interfaces.AuthServiceInterface) interfaces.AuthHandlerInterface {
	return &AuthHandler{
		service: service,
	}
}

func (handler *AuthHandler) Register(ctx *fiber.Ctx) error {
	req := new(dto.RegisterDTO)
	if err := ctx.BodyParser(req); err != nil {
		return utils.JsonError(ctx, err, "ERR_BAD_REQUEST")
	}
	return handler.service.Register(ctx, req)
}

func (handler *AuthHandler) Login(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
