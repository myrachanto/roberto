package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/myrachanto/roberto/src/synca/services"
)

// syncaController ...
var (
	SyncaController SyncaControllerInterface = &syncaController{}
)

type SyncaControllerInterface interface {
	TriggerSync(c *gin.Context)
	GetAll(c *gin.Context)
}

type syncaController struct {
	service services.SyncaServiceInterface
}

func NewsyncaController(ser services.SyncaServiceInterface) SyncaControllerInterface {
	return &syncaController{
		ser,
	}
}
func (controller syncaController) TriggerSync(c *gin.Context) {
	id := ""
	controller.service.TriggerSync(id)
	// return c.JSON(http.StatusOK, "success")
	c.JSON(http.StatusOK, gin.H{"status": "success"})
	// return
}
func (controller syncaController) GetAll(c *gin.Context) {
	syncs, err := controller.service.GetAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "failed to get logs"})
		// return c.JSON(err.Code(), err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"items": syncs})
	// return c.JSON(http.StatusOK, syncs)
	// return
}
