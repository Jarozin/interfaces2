package interfaces

import (
	"github.com/Jarozin/models"
)

type IRepoActors interface {
	GetActors() ([]*models.Actors, error)
	GetActorById(id int) (*models.Actors, error)
	CreateActor(actor *models.Actors) error
	UpdateActor(actor *models.Actors) error
	DeleteActor(id int) error
	CheckActor(actor *models.Actors) bool
}
