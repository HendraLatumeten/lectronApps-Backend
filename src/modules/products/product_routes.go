package products

import (
	"github.com/HendraLatumeten/lectronApss-Backend/src/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(rt *gin.Engine, db *gorm.DB) {
	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route := rt.Group("/products")
	{
		route.POST("", middleware.CheckAuth(), middleware.CheckAuthor(), middleware.Cloudinary(), ctrl.AddProduct)
		route.PUT("/:id", middleware.CheckAuth(), middleware.CheckAuthor(), middleware.Cloudinary(), ctrl.UpdateProduct)
		route.DELETE("/:id", middleware.CheckAuth(), middleware.CheckAuthor(), ctrl.DeleteProduct)
		route.GET("/:id", ctrl.GetById)
		route.GET("", ctrl.GetAllProduct)
		route.GET("/types/:type", ctrl.GetByType)
		route.POST("/review", middleware.CheckAuth(), ctrl.PostReview)
		route.GET("/review/:id", ctrl.GetReview)
	}

}
