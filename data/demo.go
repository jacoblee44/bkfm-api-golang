package data

import (
	"api/config"
	"api/models"
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

func GenerateDemo(db *gorm.DB, config config.Config) {
	for i := 1; i < 4; i++ {
		school := models.School{
			Name:   fmt.Sprintf("Akademie %v", i),
			Active: true,
			Domain: fmt.Sprintf("akademie%v.%v", i, config.Domain),
			Email:  fmt.Sprintf("akademie%v.test@iaccam.com", i),
			Logo:   "https://placehold.co/512x256",
			Color:  generateRandomColor(),
		}
		db.Create(&school)
		for a := 1; a < 3; a++ {
			user := models.User{
				Name:     fmt.Sprintf("Manager%v Akademie%v", a, i),
				Active:   true,
				Email:    fmt.Sprintf("manager%v.akademie%v.test@iaccam.com", a, i),
				Password: "test1234",
				Avatar:   "https://placehold.co/256x256",
			}
			db.Create(&user)
			schoolUser := models.SchoolUser{
				SchoolId: school.Id,
				UserId:   user.Id,
				// Manager:  true,
			}
			db.Create(&schoolUser)
		}
		for a := 1; a < 11; a++ {
			// module_date := getRandomTime()
			user := models.User{
				Name:     fmt.Sprintf("Fahrer%v Akademie%v", a, i),
				Active:   true,
				Email:    fmt.Sprintf("fahrer%v.akademie%v.test@iaccam.com", a, i),
				Password: "test1234",
				Avatar:   "https://placehold.co/256x256",
			}
			db.Create(&user)
			schoolUser := models.SchoolUser{
				SchoolId: school.Id,
				UserId:   user.Id,
				// Manager:  false,
				Module95: "",
			}
			db.Create(&schoolUser)
		}
	}
}

func generateRandomColor() string {
	return fmt.Sprintf("#%06x", rand.Intn(0xFFFFFF+1))
}

func getRandomTime() time.Time {

	min := time.Now().Unix()                               // Now
	max := time.Now().Add(5 * 365 * 24 * time.Hour).Unix() // 5 years in the future
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
