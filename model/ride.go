package model

import (
	"time"
)

type Ride struct {
	UUID        string    `json:"id"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	PassengerID string    `json:"-"`
	DriverID    string    `json:"-"`
	Lat         float64   `json:"lat"`
	Lon         float64   `json:"log"`
	Addr        string    `json:"address"`
	Accepted    bool      `json:"-"`
}
