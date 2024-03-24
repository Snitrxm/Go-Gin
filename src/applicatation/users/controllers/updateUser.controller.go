package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"instaclone/src/applicatation/users/dtos"
	"instaclone/src/applicatation/users/repository"
	"instaclone/src/applicatation/users/services"
	"net/http"
)

func UpdateUserController(ctx *gin.Context, userRepository repository.UserRepository) {
	userId := ctx.Param("id")

	userUUID, err := uuid.Parse(userId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	type Body struct {
		Name  *string `json:"name"`
		Email *string `json:"email"`
	}

	var body Body

	if err := ctx.ShouldBindJSON(&body); err != nil {
		return
	}

	user, err := services.UpdateUserService(dtos.UpdateUserBodyDTO{UserId: userUUID, Name: body.Name, Email: body.Email}, userRepository)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
