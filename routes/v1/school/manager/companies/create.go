package companies

import (
	"api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type CreateBody struct {
	Name           string `json:"name"`
	ContactPerson  string `json:"contact_person"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	CustomerNumber string `json:"customer_number"`
	SelfEmployed   bool   `json:"self_employed"`
	Street         string `json:"street"`
	ZipCode        string `json:"zip_code"`
	City           string `json:"city"`
	Notes          string `json:"notes"`
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

	schoolCompany := models.SchoolCompany{
		SchoolId:       school.Id,
		Name:           createBody.Name,
		ContactPerson:  createBody.ContactPerson,
		CustomerNumber: createBody.CustomerNumber,
		SelfEmployed:   createBody.SelfEmployed,
		Email:          createBody.Email,
		PhoneNumber:    createBody.PhoneNumber,
		Street:         createBody.Street,
		ZipCode:        createBody.ZipCode,
		City:           createBody.City,
		Note:           createBody.Notes,
	}

	db.Create(&schoolCompany)

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   schoolCompany,
	})
}
