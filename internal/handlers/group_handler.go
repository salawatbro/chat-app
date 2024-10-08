package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/salawatbro/chat-app/internal/dto"
	"github.com/salawatbro/chat-app/internal/interfaces"
	"github.com/salawatbro/chat-app/pkg/utils"
)

func NewGroupHandler(service interfaces.GroupServiceInterface) interfaces.GroupHandlerInterface {
	return &GroupHandler{
		service: service,
	}
}

type GroupHandler struct {
	service interfaces.GroupServiceInterface
}

func (handler *GroupHandler) FindAll(ctx *fiber.Ctx) error {
	return handler.service.FindAll(ctx)
}

func (handler *GroupHandler) FindByID(ctx *fiber.Ctx) error {
	id, err := utils.ValidateIdParams(ctx)
	if err != nil {
		return err
	}
	return handler.service.FindByID(ctx, id)
}

func (handler *GroupHandler) Create(ctx *fiber.Ctx) error {
	req := new(dto.GroupDTO)
	if err := ctx.BodyParser(req); err != nil {
		return utils.JsonError(ctx, err, "ERR_VALIDATION")
	}
	return handler.service.Create(ctx, req)
}

func (handler *GroupHandler) Update(ctx *fiber.Ctx) error {
	id, err := utils.ValidateIdParams(ctx)
	if err != nil {
		return err
	}
	req := new(dto.GroupDTO)
	if err := ctx.BodyParser(req); err != nil {
		return utils.JsonError(ctx, err, "ERR_VALIDATION")
	}
	return handler.service.Update(ctx, id, req)
}

func (handler *GroupHandler) Delete(ctx *fiber.Ctx) error {
	id, err := utils.ValidateIdParams(ctx)
	if err != nil {
		return err
	}
	return handler.service.Delete(ctx, id)
}
