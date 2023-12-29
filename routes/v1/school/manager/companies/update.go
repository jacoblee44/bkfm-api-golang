package companies

import (
	"api/models"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UpateBody struct {
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

func Update(ctx *gin.Context) {
	param := ctx.Param("id")
	var updateBody UpateBody
	log.Println("company_update_body:", updateBody, param)
	if err := ctx.ShouldBindJSON(&updateBody); err != nil {
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

	db.Model(&models.SchoolCompany{}).Where("id = ?", param).Updates(&updateBody)

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   updateBody,
	})
}
