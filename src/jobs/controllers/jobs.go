package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myrachanto/roberto/src/jobs/models"
	"github.com/myrachanto/roberto/src/jobs/services"
)

// jobController ...
var (
	JobController JobControllerInterface = jobController{}
)

type JobControllerInterface interface {
	CreateJob(c *gin.Context)
	GetJobs(c *gin.Context)
}

type jobController struct {
	service services.JobServiceInterface
}

func NewjobController(ser services.JobServiceInterface) JobControllerInterface {
	return &jobController{
		ser,
	}
}
func (controller jobController) CreateJob(c *gin.Context) {
	job := &models.Job{}
	res, err := controller.service.CreateJob(job)
	if err != nil {

		c.JSON(http.StatusOK, gin.H{"status": err.Message()})
	}
	// return c.JSON(http.StatusOK, "success")
	c.JSON(http.StatusOK, gin.H{"item": res})
	// return
}
func (controller jobController) GetJobs(c *gin.Context) {
	syncs, err := controller.service.GetJobs()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "failed to get logs"})
		// return c.JSON(err.Code(), err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": syncs})
	// return c.JSON(http.StatusOK, syncs)
	// return
}
