package users

import "github.com/RonanConway/eventsRestAPI/models"

type UserService interface {
	Save(*models.User) error
	ValidateCredentials(*models.User) error
	GetAllUsers() ([]models.User, error)
}

type userServiceImpl struct{}

func (s userServiceImpl) Save(u *models.User) error {
	return u.Save()
}

func (s userServiceImpl) ValidateCredentials(u *models.User) error {
	return u.ValidateCredentials()
}

func (s userServiceImpl) GetAllUsers() ([]models.User, error) {
	return models.GetAllUsers()
}
