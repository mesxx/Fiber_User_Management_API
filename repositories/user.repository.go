package repositories

import (
	"github.com/mesxx/Fiber_User_Management_API/models"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Create(user *models.User) (*models.User, error)
		GetAll() ([]models.User, error)
		GetByID(id uint) (*models.User, error)
		GetByEmail(email string) (*models.User, error)
		UpdateByID(user *models.User) (*models.User, error)
		DeleteByID(user *models.User) (*models.User, error)
		DeleteAll() error
	}

	userRepository struct {
		DB *gorm.DB
	}
)

func NewUserRepositoy(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (repository userRepository) Create(user *models.User) (*models.User, error) {
	if err := repository.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repository userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	if err := repository.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repository userRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	if err := repository.DB.Where("ID = ?", id).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := repository.DB.Where("email = ?", email).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository userRepository) UpdateByID(user *models.User) (*models.User, error) {
	if err := repository.DB.Where("ID = ?", user.ID).Updates(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repository userRepository) DeleteByID(user *models.User) (*models.User, error) {
	if err := repository.DB.Where("ID = ?", user.ID).Delete(user, user.ID).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repository userRepository) DeleteAll() error {
	if err := repository.DB.Where("1 = 1").Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}
