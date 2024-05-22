package ggen

import "context"

type Ggen struct {
	ID string `json:"id"`
}

func NewGgeen(ID string) Ggen {
	return Ggen{ID: ID}
}

type GgenService interface {
	GetGgen(ctx context.Context) (ggen Ggen, err error)
}

type GgenRepository interface {
	FindByID(ctx context.Context) (id string, err error)
}
