package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                 uuid.UUID `json:"id" gorm:"primaryKey"`
	Name               string    `json:"name"`
	Email              string    `json:"email"`
	Password           string    `json:"password"`
	IsAdmin            bool      `json:"is_admin"`
	FavoriteCategories string    `json:"favorite_categories"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
