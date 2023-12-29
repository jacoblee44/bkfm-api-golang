package modules

import (
	"api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func List(ctx *gin.Context) {

	db := ctx.MustGet("db").(*gorm.DB)

	var modules []models.Module

	tx := db.Model(&models.Module{}).Order("name").Find(&modules)

	if tx.Error != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"status": http.StatusForbidden,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   modules,
	})
}
