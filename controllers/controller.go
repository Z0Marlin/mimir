package controllers

import (
	"fmt"
	"net/http"

	"github.com/Z0marlin/wocbackend/models"
	"github.com/Z0marlin/wocbackend/views"
	"github.com/gin-gonic/gin"
)

var userModel models.UserModel
var userView views.UserView

func handleError(c *gin.Context, err error) {
	fmt.Println(err)
	switch {
	case err.Error() == "record not found":
		c.AbortWithStatus(http.StatusNotFound)
	default:
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
