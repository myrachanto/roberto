package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	jobcontroller "github.com/myrachanto/roberto/src/jobs/controllers"
	jobrepo "github.com/myrachanto/roberto/src/jobs/repository"
	jobservice "github.com/myrachanto/roberto/src/jobs/services"
	syncacontroller "github.com/myrachanto/roberto/src/synca/controllers"
	syncarepo "github.com/myrachanto/roberto/src/synca/repository"
	syncaservice "github.com/myrachanto/roberto/src/synca/services"
)

func ApiLoader() {
	job := jobcontroller.NewjobController(jobservice.NewjobService(jobrepo.NewjobRepository()))
	synca := syncacontroller.NewsyncaController(syncaservice.NewsyncaService(syncarepo.NewSyncaRepo()))
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	api := router.Group("/api")

	api.POST("/job", job.CreateJob)
	api.GET("/job", job.GetJobs)
	api.POST("/triggerSync", synca.TriggerSync)
	api.GET("/getsyncs", synca.GetAll)

	router.Run(":3500")
}
