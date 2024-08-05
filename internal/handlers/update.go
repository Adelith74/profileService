package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// Update profile godoc
//
//	@Summary		Update profile
//	@Description	Update profile info
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"id"
//	@Success		200	{object}	string
//	@Failure		400	{object}	string
//	@Router			/update_profile [put]
func (controller *Controller) UpdateProfile(c *gin.Context) {
	id, err := strconv.Atoi(c.Request.URL.Query().Get("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Unable to process id")
		return
	}
	fmt.Println(id)
}
