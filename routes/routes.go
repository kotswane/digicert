package routes

import (
	"digicert/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, ctrl controller.BookController) {
	r.Use(gin.Recovery())
	v1 := r.Group("/api/v1")
	{
		books := v1.Group("/books")
		{
			books.POST("/add", ctrl.Create)
			books.GET("/read", ctrl.Read)
			books.GET("/getbyid/:id", ctrl.GetById)
			books.PUT("/update/:id", ctrl.Update)
			books.DELETE("/delete/:id", ctrl.Delete)
		}
	}
}
