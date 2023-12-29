package auth

import (
	"api/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func PostLogin(ctx *gin.Context) {
	var request loginRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
		})
		return
	}
	db := ctx.MustGet("db").(*gorm.DB)
	// setup := ctx.MustGet("config").(*config.Config)
	domain := strings.Split(ctx.Request.Host, ":")[0]
	log.Println("login request:", domain)
	var school *models.School

	db.Model(&models.School{}).Where("domain = ?", domain).First(&school)

	if school == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
		})
		return
	}

	var user *models.User

	tx := db.Model(&models.User{}).Where("active = true AND email = ? AND password = ?", request.Email, request.Password).First(&user)

	if tx.Error != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status": http.StatusForbidden,
		})
		return
	}
	// var schoolUser *models.SchoolUser

	// tx = db.Model(&models.SchoolUser{}).Preload("User").Where("school_id = ? AND user_id = ?", school.Id, user.Id).First(&schoolUser)

	// if tx.Error != nil {
	// 	ctx.JSON(http.StatusForbidden, gin.H{
	// 		"status": http.StatusForbidden,
	// 	})
	// 	return
	// }

	jwt, err := GenerateJWT(user.Id, "SecretYouShouldHide")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
		})
		log.Errorln(err)
		return
	}

	ctx.JSON(200, gin.H{
		"data":    user,
		"manager": user.Manager,
		"token":   jwt,
	})
}
