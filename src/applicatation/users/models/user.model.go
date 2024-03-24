package models

import (
	"github.com/google/uuid"
	"instaclone/src/applicatation/posts/models"
	"time"
)

type UserModel struct {
	Id        uuid.UUID          `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
	Posts     []models.PostModel `gorm:"foreignKey:UserId;references:Id" json:"posts"`
}
