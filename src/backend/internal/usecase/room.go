package usecase

import (
	"context"

	"github.com/Roongjin/ChatApplication/src/backend/internal/model"
	"github.com/Roongjin/ChatApplication/src/backend/internal/repository"
	"github.com/Roongjin/ChatApplication/src/backend/internal/repository/postgres"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RoomUseCase struct {
	RoomRepo repository.Room
	LinkRepo repository.Link
	UserRepo repository.User
}

func NewRoomUseCase(db *bun.DB) *RoomUseCase {
	return &RoomUseCase{
		RoomRepo: postgres.NewRoomDB(db),
		LinkRepo: postgres.NewLinkDB(db),
		UserRepo: postgres.NewUserDB(db),
	}
}

func (r *RoomUseCase) InitBroadcastRoom(ctx context.Context) error {
	exist, err := r.RoomRepo.CheckBroadcastRoomExistence(ctx)
	if err != nil {
		return err
	}

	if !exist {
		if err := r.RoomRepo.CreateBroadcastRoom(ctx); err != nil {
			return err
		}
	}

	bcstRoom, err := r.RoomRepo.GetBroadcastRoom(ctx)
	if err != nil {
		return err
	}

	model.BroadcastRoomId = bcstRoom.Id
	return nil
}

func (r *RoomUseCase) PopulateUserInRoom(ctx context.Context, room *model.Room) error {
	links, err := r.LinkRepo.FindByRoomId(ctx, room.Id)
	if err != nil {
		return err
	}

	userIds := []uuid.UUID{}
	for _, link := range links {
		userIds = append(userIds, link.UserId)
	}

	users, err := r.UserRepo.FindByIds(ctx, userIds...)
	if err != nil {
		return err
	}

	room.Members = users
	return nil
}

func (r *RoomUseCase) GetRoomsByUserId(ctx context.Context, userId uuid.UUID) ([]*model.Room, error) {
	links, err := r.LinkRepo.FindByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	roomIds := []uuid.UUID{}
	for _, link := range links {
		if link.RoomId != model.BroadcastRoomId {
			roomIds = append(roomIds, link.RoomId)
		}
	}

	rooms, err := r.RoomRepo.FindByIds(ctx, roomIds...)
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (r *RoomUseCase) MustFindByMemberNames(ctx context.Context, memberNames []string) (*model.Room, error) {
	members, err := r.UserRepo.FindByNames(ctx, memberNames...)
	if err != nil {
		return nil, err
	}

	memberIds := []uuid.UUID{}
	for _, member := range members {
		memberIds = append(memberIds, member.Id)
	}

	exist, err := r.LinkRepo.CheckExistenceByRoomMemberIds(ctx, memberIds)
	if err != nil {
		return nil, err
	}

	if exist {
		room, err := r.LinkRepo.FindRoomByRoomMemberIds(ctx, memberIds)
		if err != nil {
			return nil, err
		}
		return room, nil
	}

	newRoom := model.Room{
		Id: uuid.New(),
	}

	if err := r.RoomRepo.AddOne(ctx, &newRoom); err != nil {
		return nil, err
	}

	newLinks := []*model.UserRoomLink{}
	for _, member := range members {
		newLinks = append(newLinks, &model.UserRoomLink{
			Id:     uuid.New(),
			UserId: member.Id,
			RoomId: newRoom.Id,
		})
	}

	if err := r.LinkRepo.AddBatch(ctx, newLinks); err != nil {
		return nil, err
	}

	return &newRoom, nil
}
