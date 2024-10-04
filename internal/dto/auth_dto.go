package dto

import "github.com/salawatbro/chat-app/pkg/utils"

type RegisterDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (dto *RegisterDTO) Validate() error {
	return utils.ExtractValidationError(dto)
}

func (dto *LoginDTO) Validate() error {
	return utils.ExtractValidationError(dto)
}
