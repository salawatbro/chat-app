package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/salawatbro/chat-app/internal/handlers"
	"github.com/salawatbro/chat-app/internal/repositories"
	"github.com/salawatbro/chat-app/internal/services"
	"github.com/salawatbro/chat-app/pkg/database"
)

func Setup(app *fiber.App, db *database.Database) {
	api := app.Group("/api")
	//repositories
	authRepo := repositories.NewAuthRepository(db.DB)

	//auth routes
	authHandler := handlers.NewAuthHandler(services.NewAuthService(authRepo))
	authRoutes := api.Group("/auth")
	authRoutes.Post("/register", authHandler.Register)
	authRoutes.Post("/login", authHandler.Login)
}
