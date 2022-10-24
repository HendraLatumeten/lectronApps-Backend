package models

import "time"

type Checkout struct {
	CheckoutId   uint      `gorm:"primaryKey" json:"id,omitempty"`
	CheckoutName string    `json:"checkout_name,omitempty"`
	UserId       string    `gorm:"type:uuid;" json:"user_id"`
	User         User      `json:"user"`
	ProductId    int       `gorm:"type:uint" json:"product_id"`
	Product      []Product `gorm:"foreignKey:ProductId;type:uint;references:ProductId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"product,omitempty"`
	DelivAddress string    `json:"deliv_address,omitempty"`
	Notes        string    `json:"notes,omitempty"`
	Discount     int       `json:"discount,omitempty"`
	Amount       int       `json:"amount,omitempty"`
	Total        int       `json:"total,omitempty"`
	CreatedAt    time.Time `gorm:"default:now(); not null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:now(); not null" json:"updated_at"`
}

type Checkouts []Checkout
