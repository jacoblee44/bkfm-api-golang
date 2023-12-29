package courses

import (
	"api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strings"
)

type AddParticipantsBody struct {
	Ids []uuid.UUID `json:"ids"`
}

func AddParticipants(ctx *gin.Context) {
	param := ctx.Param("id")

	var addParticipantsBody AddParticipantsBody
	if err := ctx.ShouldBindJSON(&addParticipantsBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := ctx.MustGet("db").(*gorm.DB)
	// setup := ctx.MustGet("config").(*config.Config)
	domain := strings.Split(ctx.Request.Host, ":")[0]

	var school *models.School

	db.Model(&models.School{}).Where("domain = ?", domain).First(&school)

	if school == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
		})
		return
	}

	var course models.Course

	tx := db.Model(&models.Course{}).Preload(clause.Associations).Where("school_id = ? AND id = ?", school.Id, param).First(&course)

	if tx.Error != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status": http.StatusForbidden,
		})
		return
	}

	participants := []models.SchoolUser{}

	for _, id := range addParticipantsBody.Ids {
		participants = append(participants, models.SchoolUser{
			Id: id,
		})
	}

	err := db.Model(&course).Association("Participants").Append(&participants)

	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status": http.StatusForbidden,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}

func RemoveParticipants(ctx *gin.Context) {
	param := ctx.Param("id")

	var addParticipantsBody AddParticipantsBody
	if err := ctx.ShouldBindJSON(&addParticipantsBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := ctx.MustGet("db").(*gorm.DB)
	// setup := ctx.MustGet("config").(*config.Config)
	domain := strings.Split(ctx.Request.Host, ":")[0]

	var school *models.School

	db.Model(&models.School{}).Where("domain = ?", domain).First(&school)

	if school == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
		})
		return
	}

	var course models.Course

	tx := db.Model(&models.Course{}).Preload(clause.Associations).Where("school_id = ? AND id = ?", school.Id, param).First(&course)

	if tx.Error != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status": http.StatusForbidden,
		})
		return
	}

	participants := []models.SchoolUser{}

	for _, id := range addParticipantsBody.Ids {
		participants = append(participants, models.SchoolUser{
			Id: id,
		})
	}

	err := db.Model(&course).Association("Participants").Delete(&participants)

	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status": http.StatusForbidden,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}

func RemoveItem(ctx *gin.Context) {
	param := ctx.Param("id")

	var addParticipantsBody AddParticipantsBody
	if err := ctx.ShouldBindJSON(&addParticipantsBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := ctx.MustGet("db").(*gorm.DB)
	// setup := ctx.MustGet("config").(*config.Config)
	domain := strings.Split(ctx.Request.Host, ":")[0]

	var school *models.School

	db.Model(&models.School{}).Where("domain = ?", domain).First(&school)

	if school == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
		})
		return
	}

	var course models.Course

	tx := db.Model(&models.Course{}).Preload(clause.Associations).Where("school_id = ? AND id = ?", school.Id, param).First(&course)

	if tx.Error != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status": http.StatusForbidden,
		})
		return
	}

	participants := []models.SchoolUser{}

	for _, id := range addParticipantsBody.Ids {
		participants = append(participants, models.SchoolUser{
			Id: id,
		})
	}

	err := db.Model(&course).Association("Participants").Delete(&participants)

	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status": http.StatusForbidden,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}
