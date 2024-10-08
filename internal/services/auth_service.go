package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/salawatbro/chat-app/internal/config"
	"github.com/salawatbro/chat-app/internal/dto"
	"github.com/salawatbro/chat-app/internal/interfaces"
	"github.com/salawatbro/chat-app/internal/middlewares"
	"github.com/salawatbro/chat-app/internal/models"
	"github.com/salawatbro/chat-app/internal/response"
	"github.com/salawatbro/chat-app/pkg/constants"
	"github.com/salawatbro/chat-app/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct {
	repo   interfaces.AuthRepositoryInterface
	Config *config.Config
}

func NewAuthService(repo interfaces.AuthRepositoryInterface, config *config.Config) interfaces.AuthServiceInterface {
	return &AuthService{
		repo:   repo,
		Config: config,
	}
}

func (service *AuthService) Register(ctx *fiber.Ctx, req *dto.RegisterDTO) error {
	if err := req.Validate(); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}
	//check if email already exists
	_, err := service.repo.FindByEmail(ctx, req.Email)
	if err == nil {
		return utils.JsonError(ctx, constants.ErrEmailExist, "ERR_EMAIL_ALREADY_EXISTS")
	}
	//hash password
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return utils.JsonError(ctx, err, "ERR_INTERNAL_SERVER")
	}
	user := models.User{
		ID:       primitive.NewObjectID(),
		Name:     req.Name,
		Email:    req.Email,
		Password: string(password),
	}
	_, err = service.repo.Register(ctx, user)
	if err != nil {
		return utils.JsonError(ctx, err, "ERR_INTERNAL_SERVER")
	}
	return utils.JsonSuccess(ctx, response.NewRegisterResponse(&user))
}

func (service *AuthService) Login(ctx *fiber.Ctx, req *dto.LoginDTO) error {
	if err := req.Validate(); err != nil {
		return utils.JsonErrorValidation(ctx, err)
	}
	user, err := service.repo.FindByEmail(ctx, req.Email)
	if err != nil {
		return utils.JsonError(ctx, constants.ErrInvalidAuth, "ERR_INVALID_AUTH")
	}
	// Compare passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return utils.JsonErrorUnauthorized(ctx, constants.ErrInvalidAuth)
	}
	// Generate token
	expireHour, _ := time.ParseDuration(service.Config.JWT.Expires.String())
	expiresAt := time.Now().Add(expireHour)
	token, err := service.generateToken(user.ID, expiresAt.Unix())
	if err != nil {
		return utils.JsonErrorUnauthorized(ctx, err)
	}
	// Save session
	session := models.Session{
		ID:     primitive.NewObjectID(),
		UserID: user.ID,
		Token:  token,
		Ip:     ctx.IP(),
		Agent:  ctx.Get("User-Agent"),
	}
	_, err = service.repo.Login(ctx, session)
	if err != nil {
		return utils.JsonError(ctx, err, "ERR_INTERNAL_SERVER")
	}
	return utils.JsonSuccess(ctx, response.NewLoginResponse(token, expiresAt))
}

func (service *AuthService) generateToken(id primitive.ObjectID, expiresAt int64) (string, error) {
	// Create JWT claims
	claims := middlewares.JwtCustomClaims{
		Issuer: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(service.Config.JWT.Secret))
}
