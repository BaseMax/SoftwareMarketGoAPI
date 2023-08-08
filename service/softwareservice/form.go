package softwareservice

import (
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type RgisterSoftwareRequest struct {
	Name        string   `json:"name" validate:"required,min=3"`
	Description string   `json:"description" validate:"max=255"`
	Price       float64  `json:"price" validate:"number"`
	Creator     string   `json:"creator" validate:"min=3"`
	Version     string   `json:"version" `
	Category    []string `json:"category_list"`
}

func (s *RgisterSoftwareRequest) Validate() error {

	v := validator.New()

	if err := v.Struct(s); err != nil {

		return fmt.Errorf("validation error : %s", err.Error())
	}

	return nil
}

type PublicFilter struct {
	Sort  string `json:"sort"`
	Limit int    `json:"limit"`
}

func MakePublicFilter(sort, limit string) (PublicFilter, error) {

	intlimit, err := strconv.Atoi(limit)
	if err != nil {

		return PublicFilter{}, fmt.Errorf("limit must be an integer")
	}

	return PublicFilter{
		Sort:  sort,
		Limit: intlimit,
	}, nil
}

type RateRequersForm struct {
	Value      uint8  `json:"value" validate:"required,gte=1,lte=5"`
	SoftwareId string `json:"software_id" `
	UserId     string `json:"user_id" `
}

func (r *RateRequersForm) Validate() error {

	v := validator.New()

	if err := v.Struct(r); err != nil {

		return fmt.Errorf("validation error : %s", err.Error())
	}

	return nil
}

type ReviewRequersForm struct {
	Title      string `json:"title" validate:"required,min=3"`
	Content    string `json:"content" validate:"required,max=255"`
	SoftwareId string `json:"software_id" `
	UserId     string `json:"user_id" `
}

func (r *ReviewRequersForm) Validate() error {

	v := validator.New()

	if err := v.Struct(r); err != nil {

		return fmt.Errorf("validation error : %s", err.Error())
	}

	return nil
}
