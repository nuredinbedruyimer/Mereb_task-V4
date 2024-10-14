package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mereb/v4/controllers"
)

func PersonRoutes(personRoutes *gin.Engine) *gin.Engine {

	personRoutes.POST("/persons", controllers.CreatePersonController)
	personRoutes.GET("/persons", controllers.GetAllPersonsController)
	personRoutes.GET("/persons/:id", controllers.GetPersonController)
	personRoutes.PUT("/persons/:id", controllers.UpdatePersonController)
	personRoutes.DELETE("/persons/:id", controllers.DeletePersonController)
	personRoutes.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Status":  "Success",
			"Message": "The requested URL " + ctx.Request.URL.Path + " was not found on this server.",
		})
	})

	return personRoutes
}
