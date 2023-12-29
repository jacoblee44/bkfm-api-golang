package v1

import (
	"api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
)

func GetTheme(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	domain := strings.Split(ctx.Request.Host, ":")[0]

	var school models.School

	db.Model(&models.School{}).Where("domain = ?", domain).First(&school)

	ctx.JSON(200, gin.H{
		"data": gin.H{
			"logo":  school.Logo,
			"color": school.Color,
		},
	})
}
