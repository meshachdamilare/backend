package handlers

import (
	"fmt"
	"quizard/helpers"
	"quizard/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserHandler struct {
	database *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		database: db,
	}
}

// UserList returns a list of users
func (u *UserHandler) UserList(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"success": true,
	})
}

// UserCreate registers a user
func (u *UserHandler) UserCreate(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
	var input helpers.InputCreateUser
	if err := c.BodyParser(&input); err != nil {
		return helpers.Dispatch500Error(c, err)
	}
	// validate email does not exist
	var user *models.User
	err := u.database.Raw("SELECT * FROM users WHERE email = ?", input.Email).Scan(&user).Error
	if err != nil {
		return helpers.Dispatch500Error(c, err)
	}
	if user != nil {
		return helpers.Dispatch400Error(c, "email already exist", nil)
	}

	hash, err := helpers.HashPassword(input.Password)
	if err != nil {
		return helpers.Dispatch500Error(c, err)
	}
	// create record
	user = &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hash,
	}
	if err := u.database.Table("users").Create(user).Error; err != nil {
		return helpers.Dispatch500Error(c, err)
	}
	return c.JSON(fiber.Map{
		"success": true,
		"message": "register user successfully",
		"data": map[string]string{
			"id":    fmt.Sprint(user.ID),
			"email": user.Email,
			"name":  user.Name,
		},
	})
}

func (u *UserHandler) GetUserById(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"user":    nil,
	})
}

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendFile("./static/private/404.html")
}
