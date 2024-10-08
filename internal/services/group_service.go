package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/salawatbro/chat-app/internal/dto"
	"github.com/salawatbro/chat-app/internal/interfaces"
	"github.com/salawatbro/chat-app/internal/models"
	"github.com/salawatbro/chat-app/internal/response"
	"github.com/salawatbro/chat-app/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type GroupService struct {
	repo interfaces.GroupRepositoryInterface
}

func NewGroupService(repo interfaces.GroupRepositoryInterface) interfaces.GroupServiceInterface {
	return &GroupService{
		repo: repo,
	}
}

func (service *GroupService) FindAll(ctx *fiber.Ctx) error {
	groups, err := service.repo.FindAll(ctx)
	if err != nil {
		return err
	}
	return utils.JsonSuccess(ctx, response.NewGroupsResponse(groups))
}

func (service *GroupService) FindByID(ctx *fiber.Ctx, id string) error {
	group, err := service.repo.FindByID(ctx, id)
	if err != nil {
		return utils.JsonError(ctx, err, "INTERNAL_SERVER_ERROR")
	}
	return utils.JsonSuccess(ctx, response.NewGroupResponse(group))
}

func (service *GroupService) Create(ctx *fiber.Ctx, req *dto.GroupDTO) error {
	if err := req.Validate(); err != nil {
		return utils.JsonError(ctx, err, "ERR_VALIDATION")
	}

	group := &models.Chat{
		ID:       primitive.NewObjectID(),
		IsGroup:  true,
		IsPublic: req.IsPublic,
		ChatName: req.Name,
		Participants: []models.Participant{
			{
				UserId:   utils.UserIdFromCtx(ctx),
				JoinedAt: primitive.NewDateTimeFromTime(time.Now()),
			},
		},
		LastMessage: nil,
		CreatedBy:   utils.UserIdFromCtx(ctx),
		CreatedAt:   primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt:   primitive.NewDateTimeFromTime(time.Now()),
	}
	err := service.repo.Create(ctx, group)
	if err != nil {
		return utils.JsonError(ctx, err, "INTERNAL_SERVER_ERROR")
	}
	return utils.JsonSuccess(ctx, response.NewGroupResponse(group))
}

func (service *GroupService) Update(ctx *fiber.Ctx, id string, req *dto.GroupDTO) error {
	if err := req.Validate(); err != nil {
		return utils.JsonError(ctx, err, "ERR_VALIDATION")
	}
	hexId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return utils.JsonError(ctx, err, "ERR_VALIDATION")
	}
	group := &models.Chat{
		ID:        hexId,
		ChatName:  req.Name,
		IsPublic:  req.IsPublic,
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}
	err = service.repo.Update(ctx, id, group)
	if err != nil {
		return utils.JsonError(ctx, err, "INTERNAL_SERVER_ERROR")
	}
	return utils.JsonSuccess(ctx, response.NewGroupResponse(group))
}

func (service *GroupService) Delete(ctx *fiber.Ctx, id string) error {
	err := service.repo.Delete(ctx, id)
	if err != nil {
		return utils.JsonError(ctx, err, "INTERNAL_SERVER_ERROR")
	}
	return utils.JsonSuccess(ctx, nil)
}
