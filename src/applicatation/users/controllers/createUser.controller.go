package controllers

import (
	"github.com/gin-gonic/gin"
	"instaclone/src/applicatation/users/dtos"
	"instaclone/src/applicatation/users/repository"
	"instaclone/src/applicatation/users/services"
	"net/http"
)

func CreateUserController(ctx *gin.Context, userRepository repository.UserRepository) {
	var body dtos.CreateUserBodyDTO

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse body"})
		return
	}

	user, err := services.CreateUserService(body, userRepository)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
