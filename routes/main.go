package routes

import (
	"api/config"
	"api/data"
	"api/db"
	"api/models"
	sentryGin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router = gin.Default()
)

func Run() {
	configData := config.Init()
	router.Use(sentryGin.New(sentryGin.Options{
		Repanic: true,
	}))

	database := db.Init(configData.Database)
	if configData.Demo {
		database.Exec("DROP SCHEMA public CASCADE")
		database.Exec("CREATE SCHEMA public")
	}
	models.Migration(database)
	if configData.Demo {
		data.GenerateDemo(database, configData)
	}
	router.Use(func(c *gin.Context) {
		c.Set("db", database)
		c.Next()
	})
	router.Use(func(c *gin.Context) {
		c.Set("config", configData)
		c.Next()
	})
	api := router.Group("/api")
	v1 := api.Group("v1")

	addV1Routes(v1)

	err := router.Run(":8085")

	if err != nil {
		log.Fatalln(err)
	}
}
