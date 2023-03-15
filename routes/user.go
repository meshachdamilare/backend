package routes

import (
	"quizard/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func registerUser(router fiber.Router, db *gorm.DB) {
	userRouter := router.Group("users")
	handler := handlers.NewUserHandler(db)

	userRouter.Get("/profile", handler.GetUserById)
}
