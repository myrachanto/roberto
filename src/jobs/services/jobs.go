package services

import (
	httperors "github.com/myrachanto/erroring"
	"github.com/myrachanto/roberto/src/jobs/models"
	"github.com/myrachanto/roberto/src/jobs/repository"
)

var (
	JobService JobServiceInterface = &jobService{}
)

type JobServiceInterface interface {
	CreateJob(*models.Job) (*models.Job, httperors.HttpErr)
	GetJobs() ([]*models.Job, httperors.HttpErr)
}
type jobService struct {
	repo repository.JobRepoInterface
}

func NewjobService(repository repository.JobRepoInterface) JobServiceInterface {
	return &jobService{
		repository,
	}
}
func (service *jobService) CreateJob(job *models.Job) (*models.Job, httperors.HttpErr) {
	return service.repo.CreateJob(job)
}
func (service *jobService) GetJobs() ([]*models.Job, httperors.HttpErr) {
	return service.repo.GetJobs()
}
