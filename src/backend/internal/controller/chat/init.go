package chat

import "github.com/uptrace/bun"

type Resolver struct{}

func NewResolver(db *bun.DB) *Resolver {
	return &Resolver{}
}
