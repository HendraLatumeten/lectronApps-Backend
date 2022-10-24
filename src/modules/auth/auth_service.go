package auth

import (
	"github.com/HendraLatumeten/lectronApss-Backend/src/database/models"
	"github.com/HendraLatumeten/lectronApss-Backend/src/interfaces"
	"github.com/HendraLatumeten/lectronApss-Backend/src/libs"
)

type auth_service struct {
	rep interfaces.AuthRepo
}
type token_response struct {
	Tokens string `json:"token"`
}

func NewService(reps interfaces.AuthRepo) *auth_service {
	return &auth_service{reps}
}

func (u auth_service) SignIn(body models.User) *libs.Response {
	checkRegist := libs.Validation(body.Email, body.Password)
	if checkRegist != nil {
		return libs.New(checkRegist.Error(), 400, true)
	}

	user, err := u.rep.FindByEmail(body.Email)
	if err != nil {
		return libs.New("email not registered, register first", 401, true)
	}
	if !libs.CheckPass(user.Password, body.Password) {
		return libs.New("wrong password", 401, true)
	}
	token := libs.NewToken(body.Email, user.Role)
	theToken, err := token.Create()
	if err != nil {
		return libs.New(err.Error(), 401, true)
	}

	return libs.New(token_response{Tokens: theToken}, 200, false)
}

func (u auth_service) SignUp(body *models.User) *libs.Response {
	check := libs.Validation(body.Email, body.Password)
	if check != nil {
		return libs.New(check.Error(), 400, true)
	}

	hassPass, err := libs.HashPassword(body.Password)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	body.Password = hassPass
	if body.Role != "admin" {
		body.Role = "user"
	}
	result, err := u.rep.Register(body)
	if err != nil {
		return libs.New(err.Error(), 401, true)
	}
	return libs.New(result, 200, false)
}
