package validators

import (
	"quizard/helpers"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

var Validator = validator.New()

func ValidateRegisterUserSchema(c *fiber.Ctx) error {
	body := new(helpers.InputCreateUser)
  c.BodyParser(&body)

	err := Validator.Struct(body)
	if err != nil {
		return helpers.SchemaError(c, err)
	}
	return c.Next()
}
