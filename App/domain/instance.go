package domain

import (
	"time"

	"github.com/google/uuid"
)

type EC2 struct {
	ID           uuid.UUID `json:"uuid" gorm:"type:uuid;primaryKey;default:uuid_generate_v4();column:id;"`
	InstanceID   *string   `json:"instance_id" gorm:"type:VARCHAR(20);column:instance_id;"`
	ImageID      *string   `json:"image_id" gorm:"type:VARCHAR(22);column:image_id;"`
	InstanceType *string   `json:"instance_type" gorm:"type:VARCHAR(25);column:instance_type;"`
	BlockDevices *string   `json:"block_devices" gorm:"type:json;column:block_devices;"`
	IPv4         *string   `json:"ipv4" gorm:"type:VARCHAR(20);column:ipv4;"`
	State        *string   `json:"state" gorm:"type:VARCHAR(13);column:state;"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null"`
}
