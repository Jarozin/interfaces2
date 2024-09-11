package interfaces

import "github.com/Jarozin/models"

type IRepoStatistic interface {
	GetStatistic() (*models.Statistic, error)
	UpdateStatistic(stat *models.Statistic) error
}
