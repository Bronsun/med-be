package routers

import (
	"gin-boilerplate/controllers"
	"gin-boilerplate/pkg/database"

	"github.com/gin-gonic/gin"
)

// NFZRoutes creates nfz endpoints
func NFZRoutes(route *gin.Engine) {
	ctrl := controllers.NFZController{DB: database.GetDB()}
	externalNfz := route.Group("/external/nfz/")
	externalNfz.GET("/save", ctrl.SaveNFZClinics)
}
