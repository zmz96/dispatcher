package driver

import (
	"dispatcher/model"
	"log"

	"gorm.io/gorm"
)

type driverRideService struct {
	logger log.Logger
	db     *gorm.DB
}

type DriverRideService interface {
	Register(driverID string) error
	GetAssignedRide(driverID string) *model.Ride
}

func NewDriverRideService(logger log.Logger, db *gorm.DB) DriverRideService {
	return &driverRideService{
		logger: logger,
		db:     db,
	}
}

func (s *driverRideService) Register(driverID string) error {
	return nil
}

func (s *driverRideService) GetAssignedRide(driverID string) *model.Ride {
	return nil
}
