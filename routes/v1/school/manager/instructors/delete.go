package instructors

import (
	"api/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strings"
)

func Delete(ctx *gin.Context) {
	param := ctx.Param("id")
	log.Println("instructor_delete_id:", param)
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

	var instructor models.Instructor

	tx := db.Model(&models.Instructor{}).Preload(clause.Associations).Where("school_id = ? AND id = ?", school.Id, param).Delete(&instructor)

	if tx.Error != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status": http.StatusForbidden,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   instructor,
	})
}
