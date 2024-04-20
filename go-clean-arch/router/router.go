package router

import (
	"github.com/gin-gonic/gin"
	"go-clean-arch/controller"
)

func UserRouter(userController controller.UserController, r *gin.Engine) {
	r.GET("/logout", userController.Logout)
	r.POST("/login", userController.Login)
	r.POST("/signup", userController.Signup)
}
