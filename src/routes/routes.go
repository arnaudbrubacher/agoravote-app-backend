package routes

import (
    "github.com/gin-gonic/gin"
    "agoravote-app-backend/src/controllers"
    "agoravote-app-backend/src/services"
)

func SetupRoutes() *gin.Engine {
    router := gin.Default()

    userService := services.NewUserService()
    userController := controllers.UserController{UserService: userService}

    groupService := services.NewGroupService()
    groupController := controllers.GroupController{GroupService: groupService}

    router.POST("/users", userController.CreateUser)
    router.POST("/users/login", userController.UserLogin)
    router.GET("/users/:id", userController.GetUser)
    router.POST("/groups", groupController.CreateGroup)
    router.GET("/groups/:id", groupController.GetGroup)
    router.POST("/votes", controllers.CreateVote)
    router.GET("/votes/:id", controllers.GetVote)

    return router
}