package usecases

import (
	"errors"

	"github.com/mesxx/Fiber_User_Management_API/models"
	"github.com/mesxx/Fiber_User_Management_API/repositories"
)

type (
	UserUsecase interface {
		Create(requestCreateUser *models.RequestCreateUser) (*models.User, error)
		GetAll() ([]models.User, error)
		GetByID(id uint) (*models.User, error)
		GetByEmail(email string) (*models.User, error)
		UpdateByID(user *models.User) (*models.User, error)
		DeleteByID(user *models.User) (*models.User, error)
		DeleteAll() error
	}

	userUsecase struct {
		UserRepository repositories.UserRepository
	}
)

func NewUserUsecase(repository repositories.UserRepository) UserUsecase {
	return &userUsecase{
		UserRepository: repository,
	}
}

func (usecase userUsecase) Create(requestCreateUser *models.RequestCreateUser) (*models.User, error) {
	user := models.User{
		Name:     requestCreateUser.Name,
		Email:    requestCreateUser.Email,
		Password: requestCreateUser.Password,
	}
	return usecase.UserRepository.Create(&user)
}

func (usecase userUsecase) GetAll() ([]models.User, error) {
	return usecase.UserRepository.GetAll()
}

func (usecase userUsecase) GetByID(id uint) (*models.User, error) {
	getByID, err := usecase.UserRepository.GetByID(id)
	if err != nil {
		return nil, err
	} else if getByID.ID == 0 {
		return nil, errors.New("user ID is invalid, please try again")
	}
	return getByID, nil
}

func (usecase userUsecase) GetByEmail(email string) (*models.User, error) {
	getByEmail, err := usecase.UserRepository.GetByEmail(email)
	if err != nil {
		return nil, err
	} else if getByEmail.ID == 0 {
		return nil, errors.New("user email is invalid, please try again")
	}
	return getByEmail, nil
}

func (usecase userUsecase) UpdateByID(user *models.User) (*models.User, error) {
	return usecase.UserRepository.UpdateByID(user)
}

func (usecase userUsecase) DeleteByID(user *models.User) (*models.User, error) {
	return usecase.UserRepository.DeleteByID(user)
}

func (usecase userUsecase) DeleteAll() error {
	return usecase.UserRepository.DeleteAll()
}
