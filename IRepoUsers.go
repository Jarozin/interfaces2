package interfaces

import (
	"github.com/Jarozin/models"
)

type IRepoUsers interface {
	GetUsers() ([]*models.Users, error)
	GetUserById(id int) (*models.Users, error)
	GetUserByLogin(login string) (*models.Users, error)
	CheckUser(login string) bool
	CreateUser(user *models.Users) error
	UpdateUser(user *models.Users) error
	DeleteUser(id int) error
}
