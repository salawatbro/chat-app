package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/salawatbro/chat-app/internal/config"
	"github.com/salawatbro/chat-app/internal/handlers"
	"github.com/salawatbro/chat-app/internal/middlewares"
	"github.com/salawatbro/chat-app/internal/repositories"
	"github.com/salawatbro/chat-app/internal/services"
	"github.com/salawatbro/chat-app/pkg/database"
)

func Setup(app *fiber.App, db *database.Database, cfg *config.Config) {
	api := app.Group("/api", middlewares.JwtMiddleware())
	//repositories
	authRepo := repositories.NewAuthRepository(db.DB)

	//auth routes
	authHandler := handlers.NewAuthHandler(services.NewAuthService(authRepo, cfg))
	authRoutes := api.Group("/auth")
	authRoutes.Post("/register", authHandler.Register)
	authRoutes.Post("/login", authHandler.Login)
}
