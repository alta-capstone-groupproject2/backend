package auth

import (
	"time"
)

type Core struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Business interface {
	Login(data Core) (token string, ID int, err error)
}

type Data interface {
	FindUser(param string) (Core, error)
}
