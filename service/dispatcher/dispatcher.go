package dispatcher

import (
	"dispatcher/model"
	"dispatcher/service/driver"
	"errors"
)

func Dispatch(ride *model.Ride) (driverID string, err error) {
	driverID = driver.Manager.GetNextAvailableDriver()
	if len(driverID) == 0 {
		err = errors.New("No Available Driver")
		return
	}
	driver.Manager.Assign(driverID, ride)
	return
}
