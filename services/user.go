package services

import (
	"github.com/CrowderSoup/drinkingaroundthe.world/database/models"
	"gorm.io/gorm"
)

type UserService interface {
	Create(*models.User) error
	GetByEmail(string) (*models.User, error)
}

type userService struct {
	db *gorm.DB
}

func (s *userService) Create(user *models.User) error {
	if err := s.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (s *userService) GetByEmail(email string) (*models.User, error) {
	var user models.User
	result := s.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
