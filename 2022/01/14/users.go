package usersapi

import (
	"context"
	"log"
	users "mario/goa/gen/users"
)

// users service example implementation.
// The example methods log the requests and return zero values.
type userssrvc struct {
	logger *log.Logger
}

// NewUsers returns the users service implementation.
func NewUsers(logger *log.Logger) users.Service {
	return &userssrvc{logger}
}

// Create implements create.
func (s *userssrvc) Create(ctx context.Context, p *users.User) (res *users.User, err error) {
	res = &users.User{}
	s.logger.Print("users.create")
	return
}

// All implements all.
func (s *userssrvc) All(ctx context.Context) (res []*users.User, err error) {
	res = []*users.User{
		{
			Name: "Mario", BirthYear: 1900,
		},
		{
			Name: "Wario", BirthYear: 2001,
		},
	}
	s.logger.Print("users.all")
	return
}
