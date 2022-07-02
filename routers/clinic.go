package routers

import (
	"gin-boilerplate/controllers"
	"gin-boilerplate/pkg/database"

	"github.com/gin-gonic/gin"
)

// ClinicRoutes creates clinic endpoints
func ClinicRoutes(route *gin.Engine) {
	ctrl := controllers.ClinicController{DB: database.GetDB()}

	// Main search handler
	main := route.Group("/clinic/")
	main.GET("/", ctrl.GetClinics)

	// Select fields search
	search := route.Group("/search/")
	search.GET("/benefit", ctrl.GetBenefits)
	search.GET("/city", ctrl.GetCity)
	search.GET("/street", ctrl.GetStreet)
	search.GET("/privateName", ctrl.GetPrivateName)

}
