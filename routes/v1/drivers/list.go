package drivers

import (
	"api/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func List(ctx *gin.Context) {
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
	log.Printf("main_driver_list: %s", tx.Statement.SQL.String())
	if tx.Error != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status": http.StatusForbidden,
		})
		return
	}

	var schoolUsers []models.SchoolUser

	db.Model(&models.SchoolUser{}).Preload("User").Where("school_id = ?", school.Id).Limit(100).Find(&schoolUsers)

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   schoolUsers,
	})
}
