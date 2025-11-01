package model

import "time"

// AircraftType defines the valid aircraft types.
type AircraftType string

const (
	ATR          AircraftType = "ATR"
	Airbus320    AircraftType = "Airbus 320"
	Boeing737Max AircraftType = "Boeing 737 Max"
)

// VoucherAssignment represents the core domain entity.
type VoucherAssignment struct {
	ID           int64
	CrewName     string
	CrewID       string
	FlightNumber string
	FlightDate   string // Using string for YYYY-MM-DD
	AircraftType AircraftType
	Seat1        string
	Seat2        string
	Seat3        string
	CreatedAt    time.Time
}

// NewVoucherAssignment creates a new assignment entity.
func NewVoucherAssignment(name, id, flightNum, date string, aircraft AircraftType, seats []string) *VoucherAssignment {
	return &VoucherAssignment{
		CrewName:     name,
		CrewID:       id,
		FlightNumber: flightNum,
		FlightDate:   date,
		AircraftType: aircraft,
		Seat1:        seats[0],
		Seat2:        seats[1],
		Seat3:        seats[2],
		CreatedAt:    time.Now().UTC(),
	}
}
