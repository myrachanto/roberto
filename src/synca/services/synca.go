package services

import (
	httperors "github.com/myrachanto/erroring"
	"github.com/myrachanto/roberto/src/synca/models"
	"github.com/myrachanto/roberto/src/synca/repository"
)

var (
	SyncaService SyncaServiceInterface = &syncaService{}
)

type SyncaServiceInterface interface {
	TriggerSync(string)
	GetAll() ([]*models.Synca, httperors.HttpErr)
}
type syncaService struct {
	repo repository.SyncaRepoInterface
}

func NewsyncaService(repository repository.SyncaRepoInterface) SyncaServiceInterface {
	return &syncaService{
		repository,
	}
}
func (service *syncaService) TriggerSync(id string) {
	service.repo.TriggerSync(id)
}
func (service *syncaService) GetAll() ([]*models.Synca, httperors.HttpErr) {
	return service.repo.GetAll()
}
