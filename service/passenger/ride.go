package passenger

import (
	"dispatcher/model"
	"dispatcher/service/dispatcher"

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

func (s *passengerRideService) NewRide(ride *model.Ride) (err error) {
	driverID, err := dispatcher.Dispatch(ride)
	if err != nil {
		return
	}
	ride.DriverID = driverID
	err = s.db.Create(ride).Error
	return
}
