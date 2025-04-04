package handler

import (
	"github.com/gin-gonic/gin"
)

type SuggestionHander interface {
	CreateSuggestion(c *gin.Context)
	UpdateSuggestion(c *gin.Context)
	ListSuggestionByFilter(c *gin.Context)
}

type Suggestion struct {
}

func NewUserHandler() SuggestionHander {
	return &Suggestion{}
}

// @Summary	CreateSuggestion
// @Tags		User
// @Accept		json
// @Produce	json
// @Param		user_id	path	string	true	"user_id"
// @Param		request		body	domain.Suggestion	true	"Suggestion info"
// @Success	200
// @Router		/suggestion [post]
func (u *Suggestion) CreateSuggestion(c *gin.Context) {

}

// @Summary	update User
// @Tags		User
// @Accept		json
// @Produce	json
// @Param		suggestion_id	path	string	true	"suggestion_id"
// @Param		request		body	domain.Suggestion	true	"Suggestion update"
// @Success	200			{object}	domain.Suggestion
// @Router		/suggestion/{suggestion_id} [patch]
func (u *Suggestion) UpdateSuggestion(c *gin.Context) {

}

// @Summary	update User
// @Tags		User
// @Accept		json
// @Produce	json
// @Param		id_book	query	int	true	"id_book"
// @Param		id_user	query	int	true	"id_user"
// @Success	200			{object}	domain.User
// @Router		/suggestion/list [get]
func (u *Suggestion) ListSuggestionByFilter(c *gin.Context) {

}
