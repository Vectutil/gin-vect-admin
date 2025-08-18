package admin

import (
	"gin-vect-admin/internal/app/handler/system"
	"gin-vect-admin/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitSystemRout(r *gin.Engine) {
	systemRoute := r.Group("admin")
	authRouter := systemRoute.Group("")
	authRouter.Use(middleware.AuthMiddleware())

	userHandler := system.NewUserHandler()

	{
		systemRoute.POST("/login", system.NewAuthHandler().Login)
		systemRoute.POST("/register", system.NewAuthHandler().Register)
	}

	userGroup := authRouter.Group("/user")
	{
		userGroup.POST("", userHandler.Create)
		userGroup.PUT("", userHandler.Update)
		userGroup.DELETE("", userHandler.Delete)
		userGroup.GET("/:id", userHandler.GetById)
		userGroup.GET("", userHandler.List)
	}

	department := authRouter.Group("/department")
	{
		h := system.NewDepartmentHandler()
		department.POST("", h.Create)
		department.PUT("", h.Update)
		department.DELETE("", h.Delete)
		department.GET("/:id", h.GetById)
		department.GET("", h.List)
		department.GET("/tree", h.GetTree)
	}

	role := authRouter.Group("/role")
	{
		h := system.NewRoleHandler()
		role.POST("", h.Create)
		role.PUT("", h.Update)
		role.DELETE("", h.Delete)
		role.GET("/:id", h.GetById)
		role.GET("", h.List)
	}
}
