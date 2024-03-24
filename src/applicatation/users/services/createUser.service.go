package services

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"instaclone/src/applicatation/users/dtos"
	"instaclone/src/applicatation/users/models"
	"instaclone/src/applicatation/users/repository"
	"time"
)

func CreateUserService(data dtos.CreateUserBodyDTO, userRepository repository.UserRepository) (*models.UserModel, error) {
	userAlreadyExists := userRepository.GetByEmail(data.Email)

	fmt.Println(userAlreadyExists)

	if userAlreadyExists != nil {
		return nil, errors.New("user already exist")
	}

	newUserId, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	userModel := models.UserModel{Id: newUserId, Name: data.Name, Email: data.Email, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	userRepository.Create(userModel)

	return &userModel, nil
}
