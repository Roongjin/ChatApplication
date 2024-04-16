package postgres

import (
	"context"

	"github.com/Roongjin/ChatApplication/src/backend/internal/model"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type LinkDB struct {
	*BaseDB[model.UserRoomLink]
}

func NewLinkDB(db *bun.DB) *LinkDB {
	type T = model.UserRoomLink

	return &LinkDB{
		BaseDB: NewBaseDB[T](db),
	}
}

func (l *LinkDB) FindByUserId(ctx context.Context, userId uuid.UUID) ([]*model.UserRoomLink, error) {
	var subscriptions []*model.UserRoomLink
	if err := l.db.NewSelect().Model(&subscriptions).Where("user_id = ?", userId).Scan(ctx, &subscriptions); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (l *LinkDB) FindByRoomId(ctx context.Context, roomId uuid.UUID) ([]*model.UserRoomLink, error) {
	var subscriptions []*model.UserRoomLink
	if err := l.db.NewSelect().Model(&subscriptions).Where("room_id = ?", roomId).Scan(ctx, &subscriptions); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (l *LinkDB) CheckExistenceByRoomMemberIds(ctx context.Context, memberIds []uuid.UUID) (bool, error) {
	var link model.UserRoomLink
	exist, err := l.db.NewSelect().Model(&link).Where("room_id != ?", model.BroadcastRoomId).Column("room_id").Group("room_id").Having("COUNT(DISTINCT CASE WHEN user_id IN (?) THEN user_id END) = ? AND COUNT(DISTINCT user_id) = ?", bun.In(memberIds), len(memberIds), len(memberIds)).Exists(ctx)
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (l *LinkDB) FindRoomByRoomMemberIds(ctx context.Context, memberIds []uuid.UUID) (*model.Room, error) {
	var link model.UserRoomLink
	if err := l.db.NewSelect().Model(&link).Where("room_id != ?", model.BroadcastRoomId).Column("room_id").Group("room_id").Having("COUNT(DISTINCT CASE WHEN user_id IN (?) THEN user_id END) = ? AND COUNT(DISTINCT user_id) = ?", bun.In(memberIds), len(memberIds), len(memberIds)).Scan(ctx, &link); err != nil {
		return nil, err
	}

	var room model.Room
	if err := l.db.NewSelect().Model(&room).Where("id = ?", link.RoomId).Scan(ctx, &room); err != nil {
		return nil, err
	}

	return &room, nil
}
