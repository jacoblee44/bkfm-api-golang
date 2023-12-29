package drivers

import (
	"api/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Dashboard(ctx *gin.Context) {
	sort := ctx.Query("s")
	user := ctx.MustGet("user").(models.User)

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

	var schoolUser models.SchoolUser

	tx := db.Model(&models.SchoolUser{}).Where("school_id = ? AND user_id = ?", school.Id, user.Id).First(&schoolUser)

	if tx.Error != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status": http.StatusForbidden,
		})
		return
	}

	var schoolUsers []models.SchoolUser

	preload := db.Model(&models.SchoolUser{}).Preload("User").Where("school_id = ?", school.Id).Limit(5)

	preload = preload.Order(sort)

	preload.Find(&schoolUsers)

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   schoolUsers,
	})
}
