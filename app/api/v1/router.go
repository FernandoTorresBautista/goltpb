package v1

import (
	"log"

	"goltpb/app/biz"

	"github.com/gin-gonic/gin"
)

// Apiv1 application structure
type Apiv1 struct {
	logger   *log.Logger
	bizLayer biz.Handle
}

// AddRoutes entrypoint to add the routes to the application
func AddRoutes(logger *log.Logger, rg *gin.RouterGroup, bizLayer *biz.Biz) {
	// create the api
	api := Apiv1{
		logger:   logger,
		bizLayer: bizLayer,
	}
	api.logger.Println("Adding routes...")

	// base line end points
	rg.GET("/ltp", api.GetInfo)
}
