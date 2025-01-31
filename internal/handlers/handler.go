package handlers

import (
	_ "github.com/Njrctr/DeNet_test/docs"
	"github.com/Njrctr/DeNet_test/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	// Swagger Документация
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	users := router.Group("/users", h.userIdentify)
	{
		id := users.Group("/:id")
		{
			id.GET("/status", h.userInfo)
			id.POST("/referrer", h.userReferrerCode)

			task := id.Group("/task")
			{
				task.POST("/complete", h.taskComplete)
			}
		}
		users.GET("/leaderboard", h.usersLeaderboard)
	}

	task := router.Group("/task")
	{
		task.POST("/create", h.taskCreate)
		task.GET("/all", h.taskGetAll)
	}

	return router

}
