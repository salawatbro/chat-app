package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/salawatbro/chat-app/internal/dto"
	"github.com/salawatbro/chat-app/internal/interfaces"
	"github.com/salawatbro/chat-app/internal/models"
	"github.com/salawatbro/chat-app/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo interfaces.AuthRepositoryInterface
}

func NewAuthService(repo interfaces.AuthRepositoryInterface) interfaces.AuthServiceInterface {
	return &AuthService{
		repo: repo,
	}
}

func (service *AuthService) Register(ctx *fiber.Ctx, req *dto.RegisterDTO) error {
	if err := req.Validate(); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return utils.JsonError(ctx, err, "ERR_INTERNAL_SERVER")
	}
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(password),
	}
	_, err = service.repo.Register(ctx, user)
	if err != nil {
		return utils.JsonError(ctx, err, "ERR_INTERNAL_SERVER")
	}
	return utils.JsonSuccess(ctx, nil)
}

func (service *AuthService) Login(ctx *fiber.Ctx, req dto.LoginDTO) error {
	//TODO implement me
	panic("implement me")
}
