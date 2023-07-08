package driver

import (
	"dispatcher/model"
	"sync"
)

type driverManager struct {
	driverRideMap    map[string]*model.Ride
	driverRideM      sync.RWMutex
	availableDriver  []string
	availableDriverM sync.Mutex
}

var Manager *driverManager = &driverManager{}

func (m *driverManager) AddAvailableDriver(id string) {
	m.availableDriverM.Lock()
	defer m.availableDriverM.Unlock()
	m.availableDriver = append(m.availableDriver, id)
}

func (m *driverManager) GetNextAvailableDriver() (id string) {
	m.availableDriverM.Lock()
	defer m.availableDriverM.Unlock()
	l := len(m.availableDriver)
	if l == 0 {
		return
	}
	id = m.availableDriver[0]
	m.availableDriver = m.availableDriver[1:]
	return
}

func (m *driverManager) Assign(driverID string, ride *model.Ride) {
	m.driverRideM.Lock()
	defer m.driverRideM.Unlock()
	m.driverRideMap[driverID] = ride
}

func (m *driverManager) GetRide(driverID string) *model.Ride {
	m.driverRideM.RLock()
	defer m.driverRideM.RUnlock()
	return m.driverRideMap[driverID]
}
