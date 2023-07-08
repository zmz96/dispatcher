package model

import "time"

type Ride struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UUID        string    `json:"uuid"`
	PassengerID string    `json:"-"`
	DriverID    string    `json:"-"`
	Lat         float64   `json:"lat"`
	Lon         float64   `json:"log"`
	Addr        string    `json:"address"`
	Accepted    bool      `json:"-"`
}
