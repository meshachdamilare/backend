package routes

import (
	"quizard/handlers"
	"quizard/validators"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func registerAuth(router fiber.Router, db *gorm.DB) {
	authRouter := router.Group("auth")
	handler := handlers.NewAuthHandler(db)

	authRouter.Post("/register", validators.ValidateRegisterUserSchema, handler.Register)
	authRouter.Get("/google/callback", handler.GoogleOauthCallback)
	authRouter.Get("/google", handler.GoogleOauth)
	authRouter.Get("/github/callback", handler.GithubOauthCallback)
	authRouter.Get("/github", handler.GithubOauth)
}
