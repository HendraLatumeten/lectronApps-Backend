package models

import "time"

type Product struct {
	ProductId   uint      `gorm:"primaryKey" json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Image       string    `json:"image,omitempty"`
	Stock       string    `json:"stock,omitempty"`
	Telphone    int       `json:"telphone,omitempty"`
	Price       int       `json:"price,omitempty"`
	Type        string    `json:"type,omitempty"`
	Sold        string    `json:"sold,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `gorm:"default:now(); not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:now(); not null" json:"updated_at"`
}

type Review struct {
	ReviewId   uint     `gorm:"primaryKey" json:"id,omitempty"`
	UserId     string   `gorm:"type:uuid" json:"user_id"`
	User       User     `json:"user"`
	ProductId  uint     `gorm:"type:uint" json:"product_id"`
	Product    Product  `json:"product,omitempty"`
	CheckoutId uint     `gorm:"type:uint" json:"checkout_id"`
	Checkout   Checkout `json:"checkout,omitempty"`
	Rate       int      `json:"rate"`
	Comment    string   `json:"comment"`
}

type Reviews []Review
type Products []Product
