package products

import (
	"github.com/HendraLatumeten/lectronApss-Backend/src/database/models"
	"github.com/HendraLatumeten/lectronApss-Backend/src/interfaces"
	"github.com/HendraLatumeten/lectronApss-Backend/src/libs"
)

type prod_servcie struct {
	prod_repo interfaces.ProductRepo
}

func NewService(reps interfaces.ProductRepo) *prod_servcie {
	return &prod_servcie{reps}
}

func (re *prod_servcie) GetAll(limit, offset int) *libs.Response {
	data, err := re.prod_repo.FindAll(limit, offset)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}
func (re *prod_servcie) GetId(id string) *libs.Response {
	data, err := re.prod_repo.FindById(id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}
func (re *prod_servcie) GetType(types string) *libs.Response {
	data, err := re.prod_repo.FindByType(types)

	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}
func (re *prod_servcie) Add(data *models.Product) *libs.Response {
	result, err := re.prod_repo.Save(data)

	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(result, 201, false)

}
func (re *prod_servcie) Update(data *models.Product, id string) *libs.Response {

	result, err := re.prod_repo.Update(data, id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(result, 201, false)
}

func (re *prod_servcie) Delete(id string) *libs.Response {
	data, err := re.prod_repo.Delete(id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *prod_servcie) PostReview(data *models.Review, email string) *libs.Response {
	user, err := re.prod_repo.GetUser(email)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	data.UserId = user.UserId
	data.User.UserId = user.UserId
	data.User.Email = email
	data.User.Name = user.Name

	res, err := re.prod_repo.GetProdId(int(data.ProductId))
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	data.ProductId = res.ProductId
	data.Product.ProductId = res.ProductId
	data.Product.Name = res.Name
	data.Product.Type = res.Type

	co, err := re.prod_repo.GetCoId(int(data.CheckoutId))
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	data.CheckoutId = co.CheckoutId
	data.Checkout.CheckoutId = co.CheckoutId
	data.Checkout.UserId = co.UserId
	data.Checkout.Amount = co.Amount
	data.Checkout.Total = co.Total

	result, err := re.prod_repo.AddReview(data)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(result, 201, false)
}

func (re *prod_servcie) GetAllReview(id int) *libs.Response {
	data, err := re.prod_repo.GetReview(id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}
