package interfaces

import (
	"github.com/HendraLatumeten/lectronApss-Backend/src/database/models"
	"github.com/HendraLatumeten/lectronApss-Backend/src/libs"
)

type ProductRepo interface {
	Save(data *models.Product) (*models.Product, error)
	Update(data *models.Product, id string) (*models.Product, error)
	Delete(id string) (*models.Product, error)

	FindAll(limit, offset int) (*models.Products, error)
	FindById(id string) (*models.Product, error)
	FindByType(types string) (*models.Products, error)
	GetProdId(id int) (*models.Product, error)
	GetUser(email string) (*models.User, error)
	GetCoId(id int) (*models.Checkout, error)
	AddReview(data *models.Review) (*models.Review, error)
	GetReview(id int) (*models.Reviews, error)
}

type ProductService interface {
	Add(data *models.Product) *libs.Response
	Update(data *models.Product, id string) *libs.Response
	Delete(id string) *libs.Response

	GetAll(limit, offset int) *libs.Response
	GetId(id string) *libs.Response
	GetType(types string) *libs.Response
	PostReview(data *models.Review, email string) *libs.Response
	GetAllReview(id int) *libs.Response
}
