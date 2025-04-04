package pay

import (
	"github.com/gin-gonic/gin"
)

type PayHandler interface {
	CreateUser(c *gin.Context)
}

type Pay struct {
}

func NewUserHandler() PayHandler {
	return &Pay{}
}

// @Summary	CreateUser
// @Tags		User
// @Accept		json
// @Produce	json
// @Param		tg_id	path	string	true	"tg_id"
// @Param		request		body	domain.User	true	"Book info"
// @Success	200
// @Router		/user [post]
func (u *Pay) Pay(c *gin.Context) {

}
