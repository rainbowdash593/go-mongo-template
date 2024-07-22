package v1

import "example/template/internal/domain/entity"

type createUserRequest struct {
	Name  string `json:"name"       binding:"required"`
	Email string `json:"email"  binding:"required"`
}

type userResponse struct {
	User *entity.User `json:"user"`
}
