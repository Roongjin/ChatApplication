package chat

import (
	"net/http"

	"github.com/Roongjin/ChatApplication/src/backend/internal/controller/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Resolver) GetConversationsByRoomId(c *gin.Context) {
	roomIdParam := c.Param("roomId")
	roomId := uuid.MustParse(roomIdParam)

	conversations, err := r.ConversationUsecase.ConversationRepo.GetConversationsByRoomId(ctx, roomId)
	if err != nil {
		util.Raise500Error(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    conversations,
	})
}
