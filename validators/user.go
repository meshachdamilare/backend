package validators

import (
	"quizard/helpers"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

var Validator = validator.New()

func ValidateCreateUserSchema(c *fiber.Ctx) error {
	body := new(helpers.InputCreateUser)
	if err := c.BodyParser(&body); err != nil {
		return helpers.SchemaError(c, err)
	}

	err := Validator.Struct(body)
	if err != nil {
		return helpers.SchemaError(c, err)
	}
	return c.Next()
}