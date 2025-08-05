package entity

import "go-projects/hexagonal-example/internal/service/entity"

type CreateUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (r CreateUserRequest) ToServiceUser() entity.User {
	return entity.User{
		Email: r.Email,
		Name:  r.Name,
	}
}
