package middleware

import (
	"api/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ManagerPermission() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("user").(models.User)

		db := ctx.MustGet("db").(*gorm.DB)

		domain := strings.Split(ctx.Request.Host, ":")[0]

		var school models.School

		tx := db.Model(&models.School{}).Where("domain = ?", domain).First(&school)
		if tx.Error != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
			})
			ctx.Abort()
		}

		// var schoolUser models.SchoolUser

		// tx = db.Model(&models.SchoolUser{}).Where("user_id = ? AND school_id = ?", user.Id, school.Id).First(&schoolUser)
		// if tx.Error != nil {
		// 	ctx.JSON(http.StatusUnauthorized, gin.H{
		// 		"status": http.StatusUnauthorized,
		// 	})
		// 	ctx.Abort()
		// }

		if user.Manager {
			ctx.Next()
			return
		}
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
		})
		ctx.Abort()
	}
}
