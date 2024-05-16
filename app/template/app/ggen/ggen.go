package ggen

type Ggen struct {
	ID string `json:"id"`
}

func NewGgeen(ID string) Ggen {
	return Ggen{ID: ID}
}

type GgenService interface {
	GetGgen() (ggen Ggen, err error)
}

type GgenRepository interface {
	FindByID() (id string, err error)
}
