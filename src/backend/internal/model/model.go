package model

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	Id            uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	Username      string    `bun:"username,type:varchar" json:"username"`
	IsOnline      bool      `bun:"is_online,type:boolean" json:"is_online"`
}

type Room struct {
	bun.BaseModel `bun:"table:rooms,alias:rooms"`
	Id            uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
}

type UserRoomLink struct {
	bun.BaseModel `bun:"table:user_room_lookup,alias:urlookup"`
	Id            uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	UserId        uuid.UUID `bun:"user_id,type:uuid" json:"user_id"`
	RoomId        uuid.UUID `bun:"room_id,type:uuid" json:"room_id"`
}

type Conversation struct {
	bun.BaseModel `bun:"table:conversations,alias:convs"`
	Id            uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	Text          string    `bun:"text,type:varchar" json:"text"`
	UserId        uuid.UUID `bun:"user_id,type:uuid" json:"user_id"`
	RoomId        uuid.UUID `bun:"room_id,type:uuid" json:"room_id"`
}
