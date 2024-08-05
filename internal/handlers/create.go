package handlers

import (
	"github.com/gin-gonic/gin"
)

// SwitchState godoc
//
//	@Summary		Switch state of a video
//	@Description	Switch by video ID. This route is used for pausing and unpausing videos from proceeding, paused goroutines wont be deleted
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"id"
//	@Success		200	{object}	int
//	@Failure		400	{object}	int
//	@Router			/switch_state [post]
func CreateProfile(c *gin.Context) {
}
