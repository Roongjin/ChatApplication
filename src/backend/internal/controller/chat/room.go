package chat

import (
	"net/http"

	"github.com/Roongjin/ChatApplication/src/backend/internal/controller/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Resolver) GetRoomsByUserId(c *gin.Context) {
	userIdParam := c.Param("userId")
	userId := uuid.MustParse(userIdParam)

	rooms, err := r.RoomUsecase.GetRoomsByUserId(ctx, userId)
	if err != nil {
		util.Raise500Error(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    rooms,
	})
}
