package models

// User model
type User struct {
	ID       uint   `gorm:"primarykey"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Waitlist struct {
	Email string `json:"email"`
	ID    uint   `gorm:"primarykey"`
}
