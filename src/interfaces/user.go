package interfaces

import (
	"github.com/HendraLatumeten/lectronApss-Backend/src/database/models"
	"github.com/HendraLatumeten/lectronApss-Backend/src/libs"
)

type UserRepo interface {
	UpdateUser(data *models.User, email string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
}

type UserService interface {
	Update(data *models.User, email string) *libs.Response
}
