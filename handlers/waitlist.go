package handlers

import (
	"fmt"
	"quizard/helpers"
	"quizard/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type WaitlistHandler struct {
	database *gorm.DB
}

type InputCreateWaitlist struct {
	Email string `json:"email" valid:"email~Invalid Email format,required~Email is required"`
}

func (u *WaitlistHandler) WaitlistCreate(c *fiber.Ctx) error {
	var input InputCreateWaitlist
	if err := c.BodyParser(&input); err != nil {
		return helpers.Dispatch500Error(c, err)
	}

	err := helpers.ValidateBody(input)
	if err != nil {
		return helpers.Dispatch500Error(c, err)
	}

	var user *models.Waitlist
	err = u.database.Raw("SELECT * FROM waitlist WHERE email = ?", input.Email).Scan(&user).Error
	if err != nil {
		return helpers.Dispatch500Error(c, err)
	}
	if user != nil {
		return helpers.Dispatch400Error(c, "email already exist", nil)
	}

	user = &models.Waitlist{
		Email: input.Email,
	}

	if err := u.database.Table("waitlist").Create(user).Error; err != nil {
		return helpers.Dispatch500Error(c, err)
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Added successfully",
		"data": map[string]string{
			"id":    fmt.Sprint(user.ID),
			"email": user.Email,
		},
	})
}

func NewWaitListHandler(db *gorm.DB) *WaitlistHandler {
	return &WaitlistHandler{
		database: db,
	}
}
