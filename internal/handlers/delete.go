package handlers

import (
	"fmt"
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
func DeleteProfile(c *gin.Context) {
	id, err := strconv.Atoi(c.Request.URL.Query().Get("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Unable to process id")
		return
	}
	fmt.Println(id)
}
