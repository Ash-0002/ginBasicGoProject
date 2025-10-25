package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourname/my-gin-app/controllers"
)

func UserRoutes(r *gin.Engine) {
	userGroup := r.Group("/users")
	{
		userGroup.GET("/", controllers.GetUsers)
		userGroup.GET("/:id", controllers.GetUserByID)
		userGroup.POST("/", controllers.CreateUser)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
	}

}
