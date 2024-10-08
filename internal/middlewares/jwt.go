package middlewares

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/salawatbro/chat-app/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"strings"
	"time"
)

type JwtCustomClaims struct {
	Issuer primitive.ObjectID `json:"issuer"`
	jwt.StandardClaims
}

type SkipperRoutesData struct {
	Method  string
	UrlPath string
}

func JwtMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Set("X-XSS-Protection", "1; mode=block")
		ctx.Set("Strict-Transport-Security", "max-age=5184000")
		ctx.Set("X-DNS-Prefetch-Control", "off")

		// skip allowlist routes
		for _, whiteList := range whiteListRoutes() {
			if ctx.Method() == whiteList.Method && ctx.Path() == whiteList.UrlPath {
				return ctx.Next()
			}
		}
		// check header token
		authorizationToken := getAuthorizationToken(ctx)
		if authorizationToken == "" {
			err := errors.New("missing Bearer token")
			return utils.JsonErrorUnauthorized(ctx, err)
		}
		// verify token
		jwtToken, err := jwt.ParseWithClaims(authorizationToken, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			return utils.JsonErrorUnauthorized(ctx, err)
		}
		claimsData := jwtToken.Claims.(*JwtCustomClaims)
		if !jwtToken.Valid {
			return utils.JsonErrorUnauthorized(ctx, errors.New("invalid token"))
		}

		// Check token expiration time
		if claimsData.ExpiresAt < time.Now().Unix() {
			return utils.JsonErrorUnauthorized(ctx, errors.New("token expired"))
		}

		if err != nil {
			return utils.JsonErrorUnauthorized(ctx, err)
		}
		utils.Logger.Info("✅ SET USER AUTH")
		ctx.Locals("user_id", claimsData.Issuer)
		return ctx.Next()
	}
}

func getAuthorizationToken(ctx *fiber.Ctx) string {
	authorizationToken := string(ctx.Request().Header.Peek("Authorization"))
	return strings.Replace(authorizationToken, "Bearer ", "", 1)
}

func whiteListRoutes() []SkipperRoutesData {
	return []SkipperRoutesData{
		{"POST", "/api/auth/register"},
		{"POST", "/api/auth/login"},
	}
}
