package models

import (
	"time"

	"github.com/google/uuid"
)

type School struct {
	Id        uuid.UUID  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name      string     `gorm:"not null,index" json:"name"`
	Domain    string     `gorm:"not null,index,unique" json:"domain"`
	Email     string     `gorm:"not null,index" json:"email"`
	Active    bool       `gorm:"default:false" json:"active"`
	Logo      string     `json:"logo"`
	Color     string     `json:"color"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}

type SchoolCompany struct {
	Id             uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	School         *School   `json:"school,omitempty"`
	SchoolId       uuid.UUID `gorm:"not null,primaryKey" json:"school_id"`
	Name           string    `gorm:"not null" json:"name"`
	ContactPerson  string    `json:"contact_person"`
	CustomerNumber string    `json:"customer_number"`
	SelfEmployed   bool      `gorm:"not null;default:false" json:"self_employed"`
	Email          string    `json:"email"`
	PhoneNumber    string    `json:"phone_number"`
	Street         string    `json:"street"`
	ZipCode        string    `json:"zip_code"`
	City           string    `json:"city"`
	Note           string    `json:"note"`
	Logo           string    `json:"logo"`
}

type SchoolUser struct {
	Id                uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	School            *School        `json:"school,omitempty"`
	SchoolId          uuid.UUID      `gorm:"not null" json:"school_id"`
	SchoolCompany     *SchoolCompany `json:"school_company,omitempty"`
	SchoolCompanyId   uuid.UUID      `json:"school_company_id"`
	SchoolCompanyName string         `json:"school_company_name"`
	User              *User          `json:"user,omitempty"`
	UserId            uuid.UUID      `gorm:"not null" json:"user_id"`
	Name              string         `gorm:"not null" json:"name"`
	DoctorType        string         `json:"doctor_type"`
	Gender            string         `json:"gender"`
	Email             string         `json:"email"`
	Birthname         string         `json:"birth_name"`
	BirthDate         string         `json:"birth_date"`
	BirthPlace        string         `json:"birth_place"`
	DrivingLicense    string         `json:"driving_license"`
	// Manager           bool           `json:"manager"`
	Module95  string     `json:"module95"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}

type Course struct {
	Id              uuid.UUID    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	School          *School      `json:"school,omitempty"`
	SchoolId        uuid.UUID    `gorm:"not null" json:"school_id"`
	Name            string       `gorm:"not null" json:"name"`
	MaxParticipants uint         `json:"max_participants"`
	Extern          bool         `json:"extern"`
	Module          Module       `json:"module"`
	ModuleId        uuid.UUID    `gorm:"not null" json:"module_id"`
	Location        Location     `json:"location"`
	LocationId      uuid.UUID    `gorm:"not null" json:"location_id"`
	Instructor      Instructor   `json:"instructor"`
	InstructorId    uuid.UUID    `gorm:"not null" json:"instructor_id"`
	Participants    []SchoolUser `gorm:"many2many:course_participants;" json:"participants"`
	StartDate       time.Time    `json:"start_date"`
	EndDate         time.Time    `json:"end_date"`
	CreatedAt       time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       *time.Time   `gorm:"index" json:"deleted_at"`
}

type Location struct {
	Id        uuid.UUID  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	School    *School    `json:"school,omitempty"`
	SchoolId  uuid.UUID  `gorm:"not null" json:"school_id"`
	Name      string     `gorm:"not null" json:"name"`
	Street    string     `json:"street"`
	ZipCode   string     `json:"zip_code"`
	City      string     `json:"city"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}

type Instructor struct {
	Id        uuid.UUID  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	School    *School    `json:"school,omitempty"`
	SchoolId  uuid.UUID  `gorm:"not null" json:"school_id"`
	Name      string     `gorm:"not null" json:"name"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}

type Module struct {
	Id   uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name string    `gorm:"not null" json:"name"`
}
