package helpers

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// 500 - internal server error
func Dispatch500Error(c *fiber.Ctx, err error) error {
	c.Status(http.StatusInternalServerError)
	return c.JSON(fiber.Map{
		"success": false,
		"message": fmt.Sprintf("%v", err),
		"data":    nil,
	})
}

// 400 - bad request
func Dispatch400Error(c *fiber.Ctx, msg string, err any) error {
	c.Status(http.StatusBadRequest)
	return c.JSON(fiber.Map{
		"success": false,
		"message": msg,
		"data":    err,
	})
}

// 404 - not found
func Dispatch404Error(c *fiber.Ctx, msg string, err any) error {
	c.Status(http.StatusNotFound)
	return c.JSON(fiber.Map{
		"success": false,
		"message": msg,
		"data":    err,
	})
}

func SchemaError(c *fiber.Ctx, err error) error {
	var errors []*IError
	for _, err := range err.(validator.ValidationErrors) {
		var el IError
		el.Field = err.Field()
		el.Tag = err.Tag()
		el.Value = err.Param()
		errors = append(errors, &el)
	}
	return Dispatch400Error(c, "invalid body schema", errors)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}