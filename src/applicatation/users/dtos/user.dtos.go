package dtos

import "github.com/google/uuid"

type CreateUserBodyDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserBodyDTO struct {
	UserId uuid.UUID
	Name   *string
	Email  *string
}
