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
	api := app.Group("/api", middlewares.JwtMiddleware(cfg))
	//repositories
	authRepo := repositories.NewAuthRepository(db.DB)
	groupRepo := repositories.NewGroupRepository(db.DB)

	//auth routes
	authHandler := handlers.NewAuthHandler(services.NewAuthService(authRepo, cfg))
	authRoutes := api.Group("/auth")
	authRoutes.Post("/register", authHandler.Register)
	authRoutes.Post("/login", authHandler.Login)
	//group routes
	groupHandler := handlers.NewGroupHandler(services.NewGroupService(groupRepo))
	groupRoutes := api.Group("/groups")
	groupRoutes.Get("/", groupHandler.FindAll)
	groupRoutes.Get("/:id", groupHandler.FindByID)
	groupRoutes.Post("/", groupHandler.Create)
	groupRoutes.Put("/:id", groupHandler.Update)
	groupRoutes.Delete("/:id", groupHandler.Delete)
}
