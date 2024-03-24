package services

import (
	"errors"
	"github.com/google/uuid"
	"instaclone/src/applicatation/users/models"
	"instaclone/src/applicatation/users/repository"
)

func GetUserByIdService(userId uuid.UUID, userRepository repository.UserRepository) (*models.UserModel, error) {
	user := userRepository.GetById(userId)

	if user == nil {
		return nil, errors.New("user not exits")
	}

	return user, nil
}
