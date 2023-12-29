package models

import (
	"github.com/google/uuid"
	"time"
)

type Organisation struct {
	Id        uuid.UUID  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name      string     `gorm:"not null;index" json:"name"`
	Email     string     `gorm:"not null;index" json:"email"`
	School    School     `json:"school"`
	SchoolId  uuid.UUID  `gorm:"not null" json:"-"`
	Active    bool       `gorm:"default:false" json:"active"`
	Logo      string     `json:"logo"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}
