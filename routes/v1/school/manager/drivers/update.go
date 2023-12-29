package drivers

import (
	"api/models"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UpdateBody struct {
	SchoolCompanyId   uuid.UUID  `json:"school_company_id"`
	SchoolCompanyName string     `json:"school_company_name"`
	UserId            uuid.UUID  `gorm:"not null" json:"user_id"`
	Name              string     `gorm:"not null" json:"name"`
	Birthname         string     `json:"birth_name"`
	BirthDate         string     `json:"birth_date"`
	BirthPlace        string     `json:"birth_place"`
	Module95          string     `json:"module95"`
	DoctorType        string     `json:"doctor_type"`
	Gender            string     `json:"gender"`
	DrivingLicense    string     `json:"driving_license"`
	Email             string     `json:"email"`
	CreatedAt         time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt         *time.Time `gorm:"index" json:"deleted_at"`
}

// type UpdateBody struct {
// 	Drivers []Driver `json:"drivers"`
// }

func Update(ctx *gin.Context) {
	param := ctx.Param("id")
	var updateBody UpdateBody
	log.Println("driverUpdateBody:", param, updateBody)
	if err := ctx.ShouldBindJSON(&updateBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println("create_bodyerror:", err.Error())
		return
	}
	log.Println("updateBody2:", updateBody)
	db := ctx.MustGet("db").(*gorm.DB)

	domain := strings.Split(ctx.Request.Host, ":")[0]

	var school *models.School

	db.Model(&models.School{}).Where("domain = ?", domain).First(&school)

	if school == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
		})
		return
	}

	db.Model(&models.SchoolUser{}).Where("id = ?", param).Updates(&updateBody)
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}
