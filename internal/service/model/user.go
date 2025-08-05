package model

import "go-projects/hexagonal-example/internal/adapter/outbound/model"

type User struct {
	ID    int
	Email string
	Name  string
}

func (r User) ToOutbondUser() model.User {
	return model.User{
		Email: r.Email,
		Name:  r.Name,
	}
}
