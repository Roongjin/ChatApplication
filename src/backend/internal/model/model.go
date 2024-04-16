package model

import (
	"time"

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
	Members       []*User   `bun:"-" json:"members"`
}

type UserRoomLink struct {
	bun.BaseModel `bun:"table:user_room_links,alias:urlinks"`
	Id            uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	UserId        uuid.UUID `bun:"user_id,type:uuid" json:"user_id"`
	RoomId        uuid.UUID `bun:"room_id,type:uuid" json:"room_id"`
}

type Conversation struct {
	bun.BaseModel `bun:"table:conversations,alias:convs"`
	Id            uuid.UUID  `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
	Text          string     `bun:"text,type:varchar" json:"text"`
	UserId        uuid.UUID  `bun:"user_id,type:uuid" json:"user_id"`
	RoomId        uuid.UUID  `bun:"room_id,type:uuid" json:"room_id"`
	CreatedAt     time.Time  `bun:"created_at,type:timestamptz,default:now()" json:"created_at"`
	UpdatedAt     time.Time  `bun:"updated_at,type:timestamptz,default:now()" json:"updated_at"`
	DeletedAt     *time.Time `bun:"deleted_at,soft_delete,nullzero,type:timestamptz" json:"deleted_at"`
}

type Message struct {
	Id         uuid.UUID `json:"id"`
	Text       string    `json:"text"`
	SenderName string    `json:"sender_name"`
	Timestamp  time.Time `json:"ts"`
}

type NewRoomInput struct {
	RoomMembersName []string `json:"room_members_name"`
}

type BroadcastRoom struct {
	bun.BaseModel `bun:"table:broadcast_room,alias:bcroom"`
	Id            uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"id"`
}

var BroadcastRoomId uuid.UUID
