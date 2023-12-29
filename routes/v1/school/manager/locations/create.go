package locations

import (
	"api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type CreateBody struct {
	Name    string `json:"name"`
	Street  string `json:"street"`
	ZipCode string `json:"zip_code"`
	City    string `json:"city"`
}

func Create(ctx *gin.Context) {

	var createBody CreateBody
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

	location := models.Location{
		SchoolId: school.Id,
		Name:     createBody.Name,
		Street:   createBody.Street,
		ZipCode:  createBody.ZipCode,
		City:     createBody.City,
	}

	db.Create(&location)

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   location,
	})
}
