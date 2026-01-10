package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid();"`

	FirstName   string `json:"first_name" gorm:"type:varchar(50);not null"`
	LastName    string `json:"last_name" gorm:"type:varchar(50);not null"`
	Email       string `json:"email" gorm:"type:varchar(150);not null;unique"`
	PhoneNumber string `json:"phone_number" gorm:"type:varchar(20);not null;unique"`

	// Local auth
	Password *string `json:"password,omitempty" gorm:"type:varchar(255)"`

	// LDAP auth
	SID               *string `json:"sid,omitempty" gorm:"type:varchar(150)"`
	CommonName        *string `json:"common_name,omitempty" gorm:"type:varchar(150)"`
	LoginName         *string `json:"login_name,omitempty" gorm:"type:varchar(150)"`
	DisplayName       *string `json:"display_name,omitempty" gorm:"type:varchar(150)"`
	DistinguishedName *string `json:"distinguished_name,omitempty" gorm:"type:varchar(255)"`

	AuthSource AuthSource `json:"auth_source" gorm:"type:smallint;not null"`

	RoleID uuid.UUID `json:"role_id" gorm:"type:uuid;not null"`

	time.Time `json:"created_at" gorm:"column:created_at;default:CURRENT_TIMESTAMP;not null"`
}

type AuthSource int16

const (
	AuthLocal AuthSource = 1
	AuthLDAP  AuthSource = 2
)
