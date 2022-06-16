package routers

import (
	"gin-boilerplate/controllers"
	"gin-boilerplate/pkg/database"

	"github.com/gin-gonic/gin"
)

func NFZRoutes(route *gin.Engine) {
	ctrl := controllers.NFZController{DB: database.GetDB()}
	v1 := route.Group("/external/nfz/")
	v1.GET("/save", ctrl.SaveNFZClinics)
}
