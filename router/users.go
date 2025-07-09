package router

import (
	"backend2/controller"

	"github.com/gin-gonic/gin"
)

func userRouter(r *gin.RouterGroup) {
	r.GET("", controller.GetAllUsers)
	r.GET("/:id", controller.GetUser)
	r.POST("", controller.CreateUser)
	r.PATCH("/:id", controller.UpdateUser)
	r.DELETE("/:id", controller.DeleteUser)
	r.GET("/sort", controller.SortUser)
}
