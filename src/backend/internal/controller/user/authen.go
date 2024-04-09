package user

import (
	"net/http"

	"github.com/Roongjin/ChatApplication/src/backend/internal/controller/util"
	"github.com/gin-gonic/gin"
)

func (r *Resolver) Authentication(c *gin.Context) {
	username := c.Param("name")

	userEntity, err := r.UserUsecase.MustGetUserByUsername(c, username)
	if err != nil {
		util.Raise500Error(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    userEntity,
	})
}
