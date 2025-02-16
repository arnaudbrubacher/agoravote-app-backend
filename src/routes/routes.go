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
    groupController := controllers.NewGroupController(groupService)

    postService := services.NewPostService()
    postController := controllers.NewPostController(postService)

    voteService := services.NewVoteService()
    voteController := controllers.NewVoteController(voteService)

    router.POST("/users", userController.CreateUser)
    router.POST("/users/login", userController.UserLogin)
    router.GET("/users/:id", userController.GetUser)
    router.POST("/groups", groupController.CreateGroup)
    router.GET("/groups/:id", groupController.GetGroup)
    router.GET("/groups", groupController.GetGroups)
    router.POST("/posts", postController.CreatePost)
    router.GET("/posts", postController.FetchPosts)
    router.POST("/votes", voteController.CreateVote)
    router.GET("/votes/:group_id", voteController.FetchVotes)

    return router
}