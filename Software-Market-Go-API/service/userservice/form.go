package userservice

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type RgisterRequest struct {
	Name     string `json:"name" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email" `
	Password string `json:"password" validate:"required,min=5"`
	IsAdmin  bool   `json:"is_admin"`
}

func (r *RgisterRequest) Validate() error {

	v := validator.New()

	if err := v.Struct(r); err != nil {

		return fmt.Errorf("validation error : %s", err.Error())
	}

	return nil
}

type LoginRequest struct {
	Name     string `json:"name" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=5"`
}

func (l *LoginRequest) Validate() error {

	v := validator.New()

	if err := v.Struct(l); err != nil {

		return fmt.Errorf("validation error : %s", err.Error())
	}

	return nil
}
