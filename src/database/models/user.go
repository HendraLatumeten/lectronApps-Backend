package models

import "time"

type User struct {
	UserId    string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();" json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty" validate:"required"`
	Role      string    `json:"role,omitempty"`
	Password  string    `json:"Password,omitempty" validate:"required"`
	Address   string    `json:"address,omitempty"`
	Phone     int       `json:"phone,omitempty"`
	Gender    string    `json:"gender,omitempty"`
	Image     string    `json:"image,omitempty"`
	Date      time.Time `json:"date,omitempty"`
	CreatedAt time.Time `gorm:"default:now(); not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:now(); not null" json:"updated_at"`
}

type Users []User
