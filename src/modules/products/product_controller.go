package products

import (
	"encoding/json"
	"strconv"

	"github.com/HendraLatumeten/lectronApss-Backend/src/database/models"
	"github.com/HendraLatumeten/lectronApss-Backend/src/interfaces"
	"github.com/HendraLatumeten/lectronApss-Backend/src/libs"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/schema"
)

type prod_ctrl struct {
	svc interfaces.ProductService
}

func NewCtrl(reps interfaces.ProductService) *prod_ctrl {
	return &prod_ctrl{svc: reps}
}

func (re *prod_ctrl) GetAllProduct(c *gin.Context) {
	v := c.Request.URL.Query().Get("pages")
	pages, _ := strconv.Atoi(v)
	limit := 12

	var offset int
	if pages == 1 {
		offset = 0
	}
	offset = limit * (pages - 1)

	re.svc.GetAll(limit, offset).Send(c)

}
func (re *prod_ctrl) GetById(c *gin.Context) {
	prodId := c.Param("id")
	re.svc.GetId(prodId).Send(c)
}

func (re *prod_ctrl) GetByType(c *gin.Context) {
	types := c.Param("type")
	re.svc.GetType(types).Send(c)

}

func (re *prod_ctrl) AddProduct(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "multipart/form-data")
	var decoder = schema.NewDecoder()
	var data models.Product

	file, exist := c.Get("image")
	if !exist {
		libs.New("claim user is not exist", 400, true)
		c.Abort()
	}

	//file upload
	image := file.(string)
	data.Image = image

	err := decoder.Decode(&data, c.Request.PostForm)
	if err != nil {
		libs.New(err.Error(), 400, true)
		c.Abort()
	}
	re.svc.Add(&data).Send(c)
}
func (re *prod_ctrl) UpdateProduct(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "multipart/form-data")

	id := c.Param("id")

	var decoder = schema.NewDecoder()
	var data models.Product

	file, exist := c.Get("image")
	if !exist {
		libs.New("file not exist", 400, true)
		c.Abort()
	}
	//file upload
	image := file.(string)
	data.Image = image

	err := decoder.Decode(&data, c.Request.PostForm)
	if err != nil {
		libs.New(err.Error(), 400, true)
		c.Abort()

	}
	re.svc.Update(&data, id).Send(c)
}

func (re *prod_ctrl) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	re.svc.Delete(id).Send(c)

}

func (re *prod_ctrl) PostReview(c *gin.Context) {
	claim_user, exist := c.Get("email")
	if !exist {
		libs.New("claim user is not exist", 400, true)
		c.Abort()
	}
	email := claim_user.(string)

	var data models.Review
	err := json.NewDecoder(c.Request.Body).Decode(&data)
	if err != nil {
		libs.New(err.Error(), 400, true)
		c.Abort()
	}
	re.svc.PostReview(&data, email).Send(c)
}

func (re *prod_ctrl) GetReview(c *gin.Context) {
	prodId := c.Param("id")
	id, _ := strconv.Atoi(prodId)

	re.svc.GetAllReview(id).Send(c)
}
