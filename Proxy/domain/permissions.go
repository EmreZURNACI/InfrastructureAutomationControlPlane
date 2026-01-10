package domain

import (
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	ID        uuid.UUID `json:"uuid" gorm:"primaryKey;type:uuid;default:uuid_generate_v4();"`
	Name      string    `json:"name" gorm:"unique;type:varchar(20);not null;"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null"`
}

//admin 	 1111
//moderator  1110
//user		 1000

// 	Read   bool      `json:"read"`
// 	Write  bool      `json:"write"`
// 	Edit   bool      `json:"edit"`
// 	Delete bool      `json:"delete"`
