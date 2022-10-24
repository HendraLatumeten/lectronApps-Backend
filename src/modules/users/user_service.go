package users

import (
	"github.com/HendraLatumeten/lectronApss-Backend/src/database/models"
	"github.com/HendraLatumeten/lectronApss-Backend/src/interfaces"
	"github.com/HendraLatumeten/lectronApss-Backend/src/libs"
)

type user_service struct {
	user_repo interfaces.UserRepo
}

func NewService(reps interfaces.UserRepo) *user_service {
	return &user_service{reps}
}

func (re *user_service) Update(data *models.User, email string) *libs.Response {
	//Get old
	oldData, err := re.user_repo.FindByEmail(email)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	valid := libs.Validation(data.Email, data.Password)
	if valid != nil {
		return libs.New(valid.Error(), 400, true)
	}

	//Hasing New Password and update data
	hassPass, err := libs.HashPassword(data.Password)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	data.Email = oldData.Email
	data.Role = oldData.Role
	data.Password = hassPass

	result, err := re.user_repo.UpdateUser(data, email)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(result, 202, false)
}
