package routes

import (
	"api/middleware"
	v1 "api/routes/v1"
	"api/routes/v1/auth"
	"api/routes/v1/drivers"
	"api/routes/v1/school"
	"api/routes/v1/school/manager/companies"

	"github.com/gin-gonic/gin"
)

func addV1Routes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("theme", v1.GetTheme)
	authGroup := routerGroup.Group("auth")
	authGroup.POST("login", auth.PostLogin)
	authGroup.GET("me", middleware.UserAuth(true), auth.Me)

	driversGroup := routerGroup.Group("drivers", middleware.UserAuth(true))
	driversGroup.GET("dashboard", drivers.Dashboard)
	driversGroup.GET("", drivers.List)
	companiesGroup := routerGroup.Group("companies", middleware.UserAuth(true))
	companiesGroup.POST("", companies.Create)
	companiesGroup.GET("", companies.List)

	userGroup := routerGroup.Group("", middleware.UserAuth(true))

	school.AddRoutes(userGroup)
}
