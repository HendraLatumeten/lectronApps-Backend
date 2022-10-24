package routers

import (
	"errors"

	"github.com/HendraLatumeten/lectronApss-Backend/src/database"
	"github.com/HendraLatumeten/lectronApss-Backend/src/modules/auth"
	"github.com/HendraLatumeten/lectronApss-Backend/src/modules/checkout"
	"github.com/HendraLatumeten/lectronApss-Backend/src/modules/products"
	"github.com/HendraLatumeten/lectronApss-Backend/src/modules/users"
	"github.com/gin-gonic/gin"
)

func New() (*gin.Engine, error) {
	mainRoute := gin.Default()

	db, err := database.New()
	if err != nil {
		return nil, errors.New("failed init database")
	}
	users.New(mainRoute, db)
	auth.New(mainRoute, db)
	products.New(mainRoute, db)
	checkout.New(mainRoute, db)

	return mainRoute, nil
}
