package chat

import (
	"net/http"

	"github.com/Roongjin/ChatApplication/src/backend/internal/controller/util"
	"github.com/Roongjin/ChatApplication/src/backend/internal/model"
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

	if err := populateUserInRooms(r, c, rooms...); err != nil {
		util.Raise500Error(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    rooms,
	})
}

func (r *Resolver) MustGetNewRoom(c *gin.Context) {
	userIdParam := c.Param("userId")
	userId := uuid.MustParse(userIdParam)
	user, err := r.UserUsecase.UserRepo.FindOneById(c, userId)
	if err != nil {
		util.Raise500Error(c, err)
		return
	}

	newRoomInput := model.NewRoomInput{}
	if err := c.BindJSON(&newRoomInput); err != nil {
		util.Raise500Error(c, err)
		return
	}

	newRoomInput.RoomMembersName = append(newRoomInput.RoomMembersName, user.Username)

	room, err := r.RoomUsecase.MustFindByMemberNames(c, newRoomInput.RoomMembersName)
	if err != nil {
		util.Raise500Error(c, err)
		return
	}

	if err := populateUserInRooms(r, c, room); err != nil {
		util.Raise500Error(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    room,
	})
}

func populateUserInRooms(r *Resolver, c *gin.Context, rooms ...*model.Room) error {
	for _, room := range rooms {
		if err := r.RoomUsecase.PopulateUserInRoom(c, room); err != nil {
			return err
		}
	}

	return nil
}
