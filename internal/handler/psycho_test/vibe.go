package psychotest

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type User struct {
}

func NewUserHandler() Handler {
	return &User{}
}

// @Summary	CreateUser
// @Tags		User
// @Accept		json
// @Produce	json
// @Param		tg_id	path	string	true	"tg_id"
// @Param		request		body	domain.User	true	"Book info"
// @Success	200
// @Router		/user [post]
func (u *User) CreateUser(c *gin.Context) {

}

// @Summary	update User
// @Tags		User
// @Accept		json
// @Produce	json
// @Param		tg_id	path	string	true	"tg_id"
// @Param		request		body	domain.User	true	"user update"
// @Success	200			{object}	domain.User
// @Router		/user/{tg_id} [patch]
func (u *User) UpdateUser(c *gin.Context) {

}
