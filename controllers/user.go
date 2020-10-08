package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
)

type UserController struct{}

type userInput struct {
	FirstName string `json:"firstName" validate:"nonzero"`
	LastName  string `json:"lastName" validate:"nonzero"`
	Password  string `json:"password" validate:"min=8"`
}

func (uc UserController) LIST(c *gin.Context) {
	users, err := userModel.ListUsers(-1)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, userView.APIUsers(&users))
}

func (uc UserController) GET(c *gin.Context) {
	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		c.Abort()
		return
	}
	usr, err := userModel.GetUserByID(c.Param("id"))
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, userView.APIUser(&usr))
}

func (uc UserController) POST(c *gin.Context) {
	var userIn userInput

	if err := c.ShouldBindJSON(&userIn); err != nil {
		handleError(c, err)
		return
	}

	if err := validator.Validate(userIn); err != nil {
		handleError(c, err)
		return
	}

	usr := userModel.NewUser(userIn.FirstName, userIn.LastName, userIn.Password)
	userModel.CreateUser(&usr)
}
