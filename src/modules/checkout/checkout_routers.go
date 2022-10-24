package checkout

import (
	"github.com/HendraLatumeten/lectronApss-Backend/src/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(rt *gin.Engine, db *gorm.DB) {
	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route := rt.Group("").Use(middleware.CheckAuth())
	{
		route.GET("/histories", ctrl.Histories)
		route.POST("/checkout", ctrl.Checkout)
	}

}
