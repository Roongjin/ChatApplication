package user

import (
	"net/http"

	"github.com/Roongjin/ChatApplication/src/backend/internal/controller/util"
	"github.com/gin-gonic/gin"
)

func (r *Resolver) GetOnlineUsers(c *gin.Context) {
	onlineUsers, err := r.UserUsecase.UserRepo.FindOnlineUsers(c)
	if err != nil {
		util.Raise500Error(c, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    onlineUsers,
	})
}
