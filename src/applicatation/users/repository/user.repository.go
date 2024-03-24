package repository

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"instaclone/src/applicatation/users/models"
)

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{database: db}
}

func (repo UserRepository) Create(user models.UserModel) {
	repo.database.Create(&user)
}

func (repo UserRepository) Update(user models.UserModel) {
	repo.database.Save(&user)
}

func (repo UserRepository) GetByEmail(email string) *models.UserModel {
	var result *models.UserModel

	if result := repo.database.Model(models.UserModel{Email: email}).First(&result); result != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
	}

	return result
}

func (repo UserRepository) GetById(id uuid.UUID) *models.UserModel {
	var result *models.UserModel

	if result := repo.database.Model(models.UserModel{Id: id}).Preload("Posts").First(&result); result != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
	}

	return result
}
