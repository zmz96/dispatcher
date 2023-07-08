package api

import (
	"dispatcher/api/middleware"
	"dispatcher/service/passenger"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log"
	"gorm.io/gorm"
)

func RegisterRouterAPI(router *gin.RouterGroup, logger log.Logger, db *gorm.DB) {
	passengerService := passenger.NewPassengerRideService(logger, db)
	passengerCtrl := NewPassengerRideCtrl(logger, passengerService)
	authMid := middleware.NewAuthMiddleware(db)
	passengerRoute := router.Group("/passenger")
	rideRoute := passengerRoute.Group("/ride")
	rideRoute.Use(authMid.AuthMiddleware())
	rideRoute.POST("/new", passengerCtrl.RequestRide)
}
