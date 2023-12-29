package models

import (
	"github.com/google/uuid"
	"time"
)

type Manager struct {
	Id        uuid.UUID  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Email     string     `gorm:"not null;unique;index" json:"email"`
	Password  string     `gorm:"not null" json:"-"`
	Active    bool       `gorm:"default:false" json:"active"`
	Avatar    string     `json:"avatar"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}
