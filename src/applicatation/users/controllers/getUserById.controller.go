package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"instaclone/src/applicatation/users/repository"
	"instaclone/src/applicatation/users/services"
	"net/http"
)

func GetUserByIdController(ctx *gin.Context, userRepository repository.UserRepository) {
	userId := ctx.Param("id")

	userUUID, err := uuid.Parse(userId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := services.GetUserByIdService(userUUID, userRepository)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
