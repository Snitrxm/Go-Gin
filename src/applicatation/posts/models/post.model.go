package models

import (
	"github.com/google/uuid"
	"time"
)

type PostModel struct {
	Id        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	UserId    uuid.UUID `json:"userId"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
