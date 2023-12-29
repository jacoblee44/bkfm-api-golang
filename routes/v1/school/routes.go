package school

import (
	"api/middleware"
	"api/routes/v1/school/manager/companies"
	"api/routes/v1/school/manager/courses"
	"api/routes/v1/school/manager/drivers"
	"api/routes/v1/school/manager/instructors"
	"api/routes/v1/school/manager/locations"
	"api/routes/v1/school/manager/modules"

	"github.com/gin-gonic/gin"
)

func AddRoutes(userGroup *gin.RouterGroup) {
	managerGroup := userGroup.Group("manager", middleware.ManagerPermission())

	managerGroup.POST("drivers", drivers.Create)
	managerGroup.GET("drivers", drivers.List)
	managerGroup.POST("drivers/update/:id", drivers.Update)
	managerGroup.GET("drivers/:id", drivers.Details)
	managerGroup.DELETE("drivers/delete/:id", drivers.Delete)

	managerGroup.POST("companies", companies.Create)
	managerGroup.GET("companies", companies.List)
	managerGroup.GET("companies/:id", companies.Details)
	managerGroup.POST("companies/update/:id", companies.Update)
	managerGroup.DELETE("companies/delete/:id", companies.Delete)

	managerGroup.POST("courses", courses.Create)
	managerGroup.GET("courses", courses.List)
	managerGroup.GET("courses/:id", courses.Details)
	managerGroup.DELETE("courses/delete/:id", courses.Delete)
	managerGroup.POST("courses/:id/participants", courses.AddParticipants)
	managerGroup.DELETE("courses/:id/participants", courses.RemoveParticipants)

	managerGroup.POST("instructors", instructors.Create)
	managerGroup.GET("instructors", instructors.List)
	managerGroup.DELETE("instructors/delete/:id", instructors.Delete)

	managerGroup.POST("locations", locations.Create)
	managerGroup.GET("locations", locations.List)
	managerGroup.DELETE("locations/delete/:id", locations.Delete)

	managerGroup.GET("modules", modules.List)
}
