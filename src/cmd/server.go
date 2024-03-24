package main

import (
	"github.com/gin-gonic/gin"
	"instaclone/src/applicatation/users/controllers"
	"instaclone/src/applicatation/users/repository"
	"instaclone/src/infra/database"
)

func main() {
	server := gin.Default()

	db := database.Initialize()

	userRepository := repository.NewUserRepository(db)

	usersRoutes := server.Group("/users")

	usersRoutes.POST("/", func(ctx *gin.Context) {
		controllers.CreateUserController(ctx, userRepository)
	})

	usersRoutes.GET("/:id", func(ctx *gin.Context) {
		controllers.GetUserByIdController(ctx, userRepository)
	})

	usersRoutes.PATCH("/:id", func(ctx *gin.Context) {
		controllers.UpdateUserController(ctx, userRepository)
	})

	err := server.Run("localhost:3000")

	if err != nil {
		return
	}
}
