package models

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	models := []interface{}{&Manager{}, &Organisation{}, &School{}, &User{}, &SchoolUser{}, &SchoolCompany{}, &Instructor{}, &Location{}, &Course{}, &Module{}}

	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			log.Fatalf("Failed to auto migrate model %T: %v", model, err)
		}
	}

}
