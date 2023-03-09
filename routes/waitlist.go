package routes

import (
	"quizard/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterWaitlist(router fiber.Router, db *gorm.DB) {
	waitlistRouter := router.Group("waitlist")

	handlers := handlers.NewWaitListHandler(db)

	waitlistRouter.Post("/signup", handlers.WaitlistCreate)
}
