package middleware

import (
	"api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func UserAuth(require bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var accessToken string

		authHeader := ctx.GetHeader("Authorization")
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) == 2 {
			accessToken = bearerToken[1]
		} else if !require {
			ctx.Next()
			return
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
			})
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			return []byte("SecretYouShouldHide"), nil
		})

		if err != nil {
			if !require {
				ctx.Next()
				return
			}
			log.Warn(err)
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
			})
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			id := uuid.MustParse(fmt.Sprintf("%v", claims["id"]))
			user := models.User{
				Id: id,
			}

			db := ctx.MustGet("db").(*gorm.DB)

			tx := db.Model(&models.User{}).First(&user)
			if tx.Error != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"status": http.StatusUnauthorized,
				})
				ctx.Abort()
			}
			ctx.Set("user", user)

			ctx.Next()

			return
		}
		if !require {
			ctx.Next()
			return
		}
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"status": http.StatusUnauthorized,
		})
		ctx.Abort()
	}
}
