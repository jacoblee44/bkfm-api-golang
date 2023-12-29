package courses

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

type CreateBody struct {
	Name            string    `json:"name"`
	InstructorId    uuid.UUID `json:"instructor_id"`
	MaxParticipants uint      `json:"max_participants"`
	Extern          bool      `json:"extern"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
	LocationId      uuid.UUID `json:"location_id"`
	ModuleId        uuid.UUID `json:"module_id"`
}

func Create(ctx *gin.Context) {

	var createBody CreateBody
	log.Println("course_create_body:", createBody)
	if err := ctx.ShouldBindJSON(&createBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
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

	course := models.Course{
		SchoolId:        school.Id,
		Name:            createBody.Name,
		InstructorId:    createBody.InstructorId,
		MaxParticipants: createBody.MaxParticipants,
		Extern:          createBody.Extern,
		LocationId:      createBody.LocationId,
		ModuleId:        createBody.ModuleId,
		StartDate:       createBody.StartDate,
		EndDate:         createBody.EndDate,
	}

	db.Create(&course)

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   course,
	})
}
