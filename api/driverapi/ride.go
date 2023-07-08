package driverapi

import (
	"dispatcher/service/driver"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log"
)

type driverRideCtrl struct {
	logger  log.Logger
	service driver.DriverRideService
}

type DriverRideCtrl interface {
	Register(ctx *gin.Context)
	GetAvailableRide(ctx *gin.Context)
}

func NewDriverRideCtrl(logger log.Logger, service driver.DriverRideService) DriverRideCtrl {
	return &driverRideCtrl{
		logger:  logger,
		service: service,
	}
}

type DriverGetReq struct {
	DriverID string `json:"id" binding:"required"`
}

func (ctrl *driverRideCtrl) Register(ctx *gin.Context) {
	// assume driver exists and is loged in
	driverReq := new(DriverGetReq)
	if err := ctx.ShouldBindJSON(&driverReq); err != nil {
		ctrl.logger.Log("failed to parse ride request [%s]", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse ride request"})
		return
	}

	ctrl.service.Register(driverReq.DriverID)
	ctx.Status(http.StatusCreated)
}

func (ctrl *driverRideCtrl) GetAvailableRide(ctx *gin.Context) {
	// assume driver exists and is loged in
	driverReq := new(DriverGetReq)
	if err := ctx.ShouldBindJSON(&driverReq); err != nil {
		ctrl.logger.Log("failed to parse ride request [%s]", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse ride request"})
		return
	}

	ride := ctrl.service.GetAssignedRide(driverReq.DriverID)
	ctx.JSON(http.StatusCreated, ride)
}
