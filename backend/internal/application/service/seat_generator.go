package service

import (
	"airline/backend/internal/domain/model"
	"fmt"
	"math/rand"
	"time"
)

// SeatGenerator is responsible for generating seat layouts and random seats.
type SeatGenerator struct {
	seatLayouts map[model.AircraftType][]string
}

// NewSeatGenerator creates a new seat generator.
func NewSeatGenerator() *SeatGenerator {
	// Pre-calculate all possible seats for each layout
	layouts := map[model.AircraftType][]string{
		model.ATR:          generateLayout(18, []string{"A", "C", "D", "F"}),
		model.Airbus320:    generateLayout(32, []string{"A", "B", "C", "D", "E", "F"}),
		model.Boeing737Max: generateLayout(32, []string{"A", "B", "C", "D", "E", "F"}),
	}
	return &SeatGenerator{seatLayouts: layouts}
}

// generateLayout is a helper to build the seat map.
func generateLayout(rows int, seats []string) []string {
	var layout []string
	for r := 1; r <= rows; r++ {
		for _, s := range seats {
			layout = append(layout, fmt.Sprintf("%d%s", r, s))
		}
	}
	return layout
}

// GenerateUniqueSeats generates 3 unique random seats for a given aircraft.
func (sg *SeatGenerator) GenerateUniqueSeats(aircraft model.AircraftType) ([]string, error) {
	layout, ok := sg.seatLayouts[aircraft]
	if !ok {
		return nil, fmt.Errorf("invalid aircraft type: %s", aircraft)
	}

	if len(layout) < 3 {
		return nil, fmt.Errorf("not enough seats available for this aircraft")
	}

	// Shuffle the layout and pick the first 3
	// We use a new source for rand to ensure it's not seeded the same every time.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create a copy to shuffle, so the original layout is not modified
	shuffledLayout := make([]string, len(layout))
	copy(shuffledLayout, layout)

	r.Shuffle(len(shuffledLayout), func(i, j int) {
		shuffledLayout[i], shuffledLayout[j] = shuffledLayout[j], shuffledLayout[i]
	})

	return shuffledLayout[:3], nil
}
