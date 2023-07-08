package driver

import (
	"dispatcher/model"

	"gorm.io/gorm"

	"github.com/go-kit/kit/log"
)

type driverRideService struct {
	logger log.Logger
	db     *gorm.DB
}

type DriverRideService interface {
	Register(driverID string)
	GetAssignedRide(driverID string) *model.Ride
}

func NewDriverRideService(logger log.Logger, db *gorm.DB) DriverRideService {
	return &driverRideService{
		logger: logger,
		db:     db,
	}
}

func (s *driverRideService) Register(driverID string) {
	Manager.AddAvailableDriver(driverID)
}

func (s *driverRideService) GetAssignedRide(driverID string) *model.Ride {
	Manager.GetRide(driverID)
	return nil
}
