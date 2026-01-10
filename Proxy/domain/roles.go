package domain

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID        uuid.UUID `json:"uuid" gorm:"primaryKey;type:uuid;default:uuid_generate_v4();"`
	Name      string    `json:"name" gorm:"unique;type:varchar(9);not null;"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null"`
}

//  Admin     string    `json:"admin"`
// 	Moderator string    `json:"moderator"`
// 	User      string    `json:"user"`
