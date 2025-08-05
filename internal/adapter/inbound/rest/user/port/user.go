package port

import "go-projects/hexagonal-example/internal/service/model"

type CreateUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (r CreateUserRequest) ToServiceUser() model.User {
	return model.User{
		Email: r.Email,
		Name:  r.Name,
	}
}
