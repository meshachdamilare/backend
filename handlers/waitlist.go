package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type WaitlistHandler struct {
	database *gorm.DB
}

func (u *WaitlistHandler) WaitlistCreate(c *fiber.Ctx) error {
	return nil
}

func NewWaitListHandler(db *gorm.DB) *WaitlistHandler {
	return &WaitlistHandler{
		database: db,
	}
}
