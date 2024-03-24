package services

import (
	"errors"
	"instaclone/src/applicatation/users/dtos"
	"instaclone/src/applicatation/users/models"
	"instaclone/src/applicatation/users/repository"
)

func UpdateUserService(data dtos.UpdateUserBodyDTO, userRepository repository.UserRepository) (*models.UserModel, error) {
	user := userRepository.GetById(data.UserId)

	if user == nil {
		return nil, errors.New("user not found")
	}

	if data.Name != nil {
		user.Name = *data.Name
	}

	if data.Email != nil {
		user.Email = *data.Email
	}

	userRepository.Update(*user)

	return user, nil
}
