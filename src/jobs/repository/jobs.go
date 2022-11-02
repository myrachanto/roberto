package repository

import (
	httperors "github.com/myrachanto/erroring"
	"github.com/myrachanto/roberto/src/jobs/models"
)

var (
	Jobrepository JobRepoInterface = &jobrepository{}
)

type JobRepoInterface interface {
	CreateJob(*models.Job) (*models.Job, httperors.HttpErr)
	GetJobs() ([]*models.Job, httperors.HttpErr)
}
type jobrepository struct {
}

func NewjobRepository() JobRepoInterface {
	return &jobrepository{}
}
func (r *jobrepository) CreateJob(*models.Job) (*models.Job, httperors.HttpErr) {
	return nil, nil
}
func (r *jobrepository) GetJobs() ([]*models.Job, httperors.HttpErr) {
	return nil, nil
}
