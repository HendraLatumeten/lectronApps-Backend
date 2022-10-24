package interfaces

import (
	"github.com/HendraLatumeten/lectronApss-Backend/src/database/models"
	"github.com/HendraLatumeten/lectronApss-Backend/src/libs"
)

type AuthRepo interface {
	FindByEmail(username string) (*models.User, error)
	Register(data *models.User) (*models.User, error)
}

type AuthService interface {
	SignIn(body models.User) *libs.Response
	SignUp(body *models.User) *libs.Response
}
