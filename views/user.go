package views

import (
	"github.com/Z0marlin/mimir/models"
	"github.com/gin-gonic/gin"
)

type UserView struct{}

func (uv UserView) APIUser(user *models.User) gin.H {
	return gin.H{
		"firstName": user.FirstName,
		"lastName":  user.LastName,
	}
}

func (uv UserView) APIUsers(users *[]models.User) []gin.H {
	ret := make([]gin.H, 0, len(*users)+1)
	for _, user := range *users {
		ret = append(ret, uv.APIUser(&user))
	}
	return ret
}
