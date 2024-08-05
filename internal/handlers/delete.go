package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// DeteleProfile godoc
//
//	@Summary		Detele profile
//	@Description	Detele profile
//	@Accept			json
//	@Produce		json
//	@Param			id	query		int	true	"id"
//	@Success		200	{object}	string
//	@Failure		400	{object}	string
//	@Router			/detele_profile [delete]
func (controller *Controller) DeleteProfile(c *gin.Context) {
	id, err := strconv.Atoi(c.Request.URL.Query().Get("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Unable to process id")
		return
	}
	fmt.Println(id)
}
