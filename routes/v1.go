package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func welcome(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"message": "quizard backend api",
		"data":    nil,
	})
}

func Routes(app *fiber.App, db *gorm.DB) {
	apiURL := "/api/v1"
	router := app.Group(apiURL)
	app.Get(apiURL, welcome)

	registerUser(router, db)
	registerAuth(router, db)
	RegisterWaitlist(router, db)
}
