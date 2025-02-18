package routes

import (
	"agoravote-app-backend/src/controllers"
	"agoravote-app-backend/src/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	groupService := services.NewGroupService()
	groupController := controllers.GroupController{GroupService: groupService}

	router.POST("/groups", groupController.CreateGroup)
	router.GET("/groups/:id", groupController.GetGroup)
	router.GET("/groups", groupController.GetGroups)

	return router
}
