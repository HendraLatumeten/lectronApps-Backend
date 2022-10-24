package interfaces

import (
	"github.com/HendraLatumeten/lectronApss-Backend/src/database/models"
	"github.com/HendraLatumeten/lectronApss-Backend/src/libs"
)

type CoRepo interface {
	FindAll() (*models.Checkouts, error)
	GetId(email string) (*models.User, error)
	GetProductId(id int) (*models.Product, error)
	FindData(id string) (*models.Checkouts, error)
	Save(data *models.Checkout) (*models.Checkout, error)
	FindCoId(id int) (*models.Checkout, error)
}

type CoService interface {
	GetAll(email string) *libs.Response
	Checkout(data *models.Checkout, email string) *libs.Response
}
