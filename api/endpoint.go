package api

import (
	"dispatcher/api/driverapi"
	"dispatcher/api/middleware"
	"dispatcher/api/passengerapi"
	"dispatcher/service/driver"
	"dispatcher/service/passenger"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log"
	"gorm.io/gorm"
)

func RegisterRouterAPI(router *gin.Engine, logger log.Logger, db *gorm.DB) {
	RegisterPassengerRoutes(router.Group("/passenger"), logger, db)
	RegisterDriverRoutes(router.Group("/driver"), logger, db)
}

func RegisterPassengerRoutes(router *gin.RouterGroup, logger log.Logger, db *gorm.DB) {
	passengerService := passenger.NewPassengerRideService(logger, db)
	passengerCtrl := passengerapi.NewPassengerRideCtrl(logger, passengerService)
	authMid := middleware.NewAuthMiddleware(db)
	rideRoute := router.Group("/ride")
	rideRoute.Use(authMid.AuthMiddleware())
	rideRoute.POST("/new", passengerCtrl.RequestRide)
}

func RegisterDriverRoutes(router *gin.RouterGroup, logger log.Logger, db *gorm.DB) {
	driverService := driver.NewDriverRideService(logger, db)
	driverCtrl := driverapi.NewDriverRideCtrl(logger, driverService)
	rideRoute := router.Group("/ride")
	rideRoute.GET("/register", driverCtrl.Register)
	rideRoute.GET("/available", driverCtrl.GetAvailableRide)
}
