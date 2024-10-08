package dto

import "github.com/salawatbro/chat-app/pkg/utils"

type GroupDTO struct {
	Name     string `json:"name" validate:"required"`
	IsPublic bool   `json:"is_public"`
}

func (dto *GroupDTO) Validate() error {
	return utils.ExtractValidationError(dto)
}
