package dtos

type UserDTO struct {
	Name     string `json:"name" validate:"required"`
	CPF      string `json:"cpf" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Type     string `json:"type" validate:"required,oneof=COMUM SELLER"`
}
