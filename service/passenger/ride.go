package passenger

import (
	"dispatcher/model"

	"github.com/go-kit/kit/log"

	"gorm.io/gorm"
)

type passengerRideService struct {
	logger log.Logger
	db     *gorm.DB
}

type PassengerRideService interface {
	NewRide(ride *model.Ride) error
}

func NewPassengerRideService(logger log.Logger, db *gorm.DB) PassengerRideService {
	return &passengerRideService{
		logger: logger,
		db:     db,
	}
}

func (s *passengerRideService) NewRide(ride *model.Ride) error {
	return s.db.Create(ride).Error
}
