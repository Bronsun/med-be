package routers

import (
	"gin-boilerplate/controllers"
	"gin-boilerplate/pkg/database"
	"github.com/gin-gonic/gin"
)

func ClinicRoutes(route *gin.Engine) {
	ctrl := controllers.ClinicController{DB: database.GetDB()}
	v1 := route.Group("/clinic")
	v1.GET("/addNew/", ctrl.AddClinics)
}
