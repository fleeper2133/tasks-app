package handler

import (
	"github.com/fleeper2133/tasks-app/internal/service"
	"github.com/gin-gonic/gin"

	_ "github.com/fleeper2133/tasks-app/docs"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/refresh", h.RefreshToken)
	}

	api := router.Group("/api")
	{
		tasks := api.Group("/tasks", h.userIdentity)
		{
			tasks.GET("/", h.GetAllTasks)
			tasks.POST("/", h.CreateTask)
			tasks.GET("/:id", h.GetTaskById)
			tasks.DELETE("/:id", h.DeleteTask)
			tasks.PUT("/:id", h.UpdateTask)
		}
	}

	return router
}
