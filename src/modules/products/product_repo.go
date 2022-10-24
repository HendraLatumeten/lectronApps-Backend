package products

import (
	"errors"
	"strings"

	"github.com/HendraLatumeten/lectronApss-Backend/src/database/models"
	"github.com/HendraLatumeten/lectronApss-Backend/src/libs"
	"gorm.io/gorm"
)

type prod_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *prod_repo {
	return &prod_repo{db}
}

func (r *prod_repo) FindAll(limit, offset int) (*models.Products, error) {
	var datas *models.Products
	result := r.db.Limit(limit).Offset(offset).Find(&datas)
	if result.Error != nil {
		return nil, errors.New("failed obtain datas")
	}
	return datas, nil
}
func (r *prod_repo) FindById(id string) (*models.Product, error) {
	var datas *models.Product

	result := r.db.First(&datas, "product_id = ?", id)
	if result.Error != nil {
		return nil, errors.New("data not found!")
	}
	return datas, nil
}
func (r *prod_repo) FindByType(types string) (*models.Products, error) {
	var datas *models.Products

	params := strings.ToLower((types))
	result := r.db.Where("LOWER(type) LIKE ?", "%"+params+"%").Find(&datas)
	if result.Error != nil {
		return nil, errors.New("data not found!")
	}
	return datas, nil
}

func (r *prod_repo) Save(data *models.Product) (*models.Product, error) {

	result := r.db.Create(data)
	if result.Error != nil {
		return nil, errors.New("failled to obtain data")
	}

	return data, nil
}
func (r *prod_repo) Update(data *models.Product, id string) (*models.Product, error) {
	var datas *models.Product

	r.db.First(&datas, "product_id = ?", id)
	if datas.Image != "" {
		libs.CloudDelete(datas.Image)
	}

	result := r.db.Model(&data).Where("product_id = ?", id).Updates(&data)
	if result.Error != nil {
		return nil, errors.New("failled to obtain data")
	}

	return data, nil
}
func (r *prod_repo) Delete(id string) (*models.Product, error) {
	var datas *models.Product

	r.db.First(&datas, "product_id = ?", id)
	if datas.Image != "" {
		libs.CloudDelete(datas.Image)
	}

	result := r.db.Where("product_id", id).Delete(&datas)
	if result.Error != nil {
		return nil, errors.New("data not found!")
	}
	return datas, nil
}

func (r *prod_repo) GetProdId(id int) (*models.Product, error) {
	var data *models.Product

	res := r.db.Where("product_id = ?", id).Find(&data)
	if res.Error != nil {
		return nil, errors.New("product id not found!")
	}
	return data, nil
}

func (r *prod_repo) GetUser(email string) (*models.User, error) {
	var data *models.User

	res := r.db.Where("email = ?", email).Find(&data)
	if res.Error != nil {
		return nil, errors.New("user email not found!")
	}
	return data, nil
}

func (r *prod_repo) GetCoId(id int) (*models.Checkout, error) {
	var data *models.Checkout

	res := r.db.Where("checkout_id = ?", id).Find(&data)
	if res.Error != nil {
		return nil, errors.New("checkout id not found!")
	}
	return data, nil
}

func (r *prod_repo) AddReview(data *models.Review) (*models.Review, error) {
	result := r.db.Create(data)
	if result.Error != nil {
		return nil, errors.New("failled to obtain data")
	}

	return data, nil
}

func (r *prod_repo) GetReview(id int) (*models.Reviews, error) {
	var datas *models.Reviews

	res := r.db.Preload("Product").Preload("Checkout").Order("review_id asc").Where("product_id = ?", id).Find(&datas)
	if res.Error != nil {
		return nil, errors.New("failled to obtain data")
	}

	return datas, nil
}
