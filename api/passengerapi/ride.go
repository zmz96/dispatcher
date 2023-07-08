package passengerapi

import (
	"dispatcher/model"
	"dispatcher/service/passenger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log"
)

type passengerRideCtrl struct {
	logger  log.Logger
	service passenger.PassengerRideService
}

type PassengerRideCtrl interface {
	RequestRide(ctx *gin.Context)
}

type RideReq struct {
	PassengerID string  `json:"id" binding:"required"`
	Lat         float64 `json:"lat" binding:"required"`
	Lon         float64 `json:"lon" binding:"required"`
	Addr        string  `json:"addr" binding:"required"`
}

func NewPassengerRideCtrl(logger log.Logger, service passenger.PassengerRideService) PassengerRideCtrl {
	return &passengerRideCtrl{
		logger:  logger,
		service: service,
	}
}

func (ctrl *passengerRideCtrl) RequestRide(ctx *gin.Context) {
	rideReq := new(RideReq)
	if err := ctx.ShouldBindJSON(&rideReq); err != nil {
		ctrl.logger.Log("failed to parse ride request [%s]", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse ride request"})
		return
	}

	ride := model.Ride{
		PassengerID: rideReq.PassengerID,
		Lon:         rideReq.Lon,
		Lat:         rideReq.Lat,
		Addr:        rideReq.Addr,
	}

	err := ctrl.service.NewRide(&ride)
	if err != nil {
		ctrl.logger.Log("failed to create credential definition [%s]", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create credential definition"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"ride_id": ride.UUID})

}
