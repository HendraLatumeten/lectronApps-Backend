package checkout

import (
	"encoding/json"

	"github.com/HendraLatumeten/lectronApss-Backend/src/database/models"
	"github.com/HendraLatumeten/lectronApss-Backend/src/interfaces"
	"github.com/HendraLatumeten/lectronApss-Backend/src/libs"
	"github.com/gin-gonic/gin"
)

type co_ctrl struct {
	svc interfaces.CoService
}

func NewCtrl(reps interfaces.CoService) *co_ctrl {
	return &co_ctrl{svc: reps}
}

func (re *co_ctrl) Histories(c *gin.Context) {
	claim_user, exist := c.Get("email")
	if !exist {
		libs.New("claim user is not exist", 400, true)
		c.Abort()
	}
	email := claim_user.(string)
	re.svc.GetAll(email).Send(c)
}

func (re *co_ctrl) Checkout(c *gin.Context) {
	claim_user, exist := c.Get("email")
	if !exist {
		libs.New("claim user is not exist", 400, true)
		c.Abort()
	}
	email := claim_user.(string)

	var data models.Checkout
	err := json.NewDecoder(c.Request.Body).Decode(&data)
	if err != nil {
		libs.New(err.Error(), 400, true)
		c.Abort()
	}
	re.svc.Checkout(&data, email).Send(c)
}
