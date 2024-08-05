package handlers

import (
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// Create profile godoc
//
//	@Summary		Create profile
//	@Description	Create profile
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"id"
//	@Success		200	{object}	string
//	@Failure		400	{object}	string
//	@Router			/create_profile [post]
func (controller *Controller) CreateProfile(c *gin.Context) {
}
