package handlers

import (
	"net/http"
	"strconv"

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
func (controller *Controller) GetProfile(c *gin.Context) {
	id, err := strconv.Atoi(c.Request.URL.Query().Get("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Unable to process id "+err.Error())
		return
	}
	profile, err := controller.db.GetProfile(id)
	if err != nil {
		c.String(http.StatusBadRequest, "")
		return
	}
	c.JSON(200, profile)
}

func (controller *Controller) GetProfiles(c *gin.Context) {
	profiles, err := controller.db.GetProfiles()
	if err != nil {
		c.String(http.StatusBadRequest, "Unable to fetch profiles: "+err.Error())
		return
	}
	c.JSON(200, profiles)
}
