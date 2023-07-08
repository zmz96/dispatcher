package dispatcher

import (
	"dispatcher/model"
	"dispatcher/service/driver"
	"errors"
)

func Dispatch(ride *model.Ride) error {
	driverID := driver.Manager.GetNextAvailableDriver()
	if len(driverID) == 0 {
		return errors.New("No Available Driver")
	}
	driver.Manager.Assign(driverID, ride)
	return nil
}
