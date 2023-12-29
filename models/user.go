package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name      string     `gorm:"not null" json:"name"`
	Email     string     `gorm:"not null,unique" json:"email"`
	Password  string     `gorm:"not null" json:"-"`
	Active    bool       `gorm:"default:false" json:"active"`
	Manager   bool       `json:manager`
	Avatar    string     `json:"avatar"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}
