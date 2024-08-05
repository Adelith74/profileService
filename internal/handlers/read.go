package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// GetProfile godoc
//
//	@Summary		Get profile
//	@Description	Get profile
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"id"
//	@Success		200	{object}	model.Profile
//	@Failure		400	{object}	string
//	@Router			/get_profile [post]
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

// @BasePath /api/v1

// GetProfiles godoc
//
//	@Summary		Get profiles
//	@Description	Get profiles
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]model.Profile
//	@Failure		400	{object}	string
//	@Router			/get_profiles [get]
func (controller *Controller) GetProfiles(c *gin.Context) {
	profiles, err := controller.db.GetProfiles()
	if err != nil {
		c.String(http.StatusBadRequest, "Unable to fetch profiles: "+err.Error())
		return
	}
	c.JSON(200, profiles)
}
