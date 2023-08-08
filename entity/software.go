package entity

import (
	"time"

	"github.com/google/uuid"
)

type (
	Software struct {
		ID          uuid.UUID `json:"id" gorm:"primaryKey"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Price       float64   `json:"price"`
		Creator     string    `json:"creator"`
		Version     string    `json:"version"`
		Category    string    `json:"category_list" `
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
	Softwares []Software

	Rate struct {
		ID         uuid.UUID `json:"id" gorm:"primaryKey"`
		Value      uint8     `json:"value"` // value can be between 1 and 5
		SoftwareId uuid.UUID `json:"software_id"`
		UserId     uuid.UUID `json:"user_id" `
		RatedAt    time.Time `json:"Rated_at"`
	}
	Rates []Rate

	Review struct {
		ID         uuid.UUID `json:"id" gorm:"primaryKey"`
		Title      string    `json:"title"`
		Content    string    `json:"content"`
		SoftwareId uuid.UUID `json:"software_id" `
		UserId     uuid.UUID `json:"user_id" `
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}
	Reviews []Review
)
