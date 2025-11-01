package service

import (
	"strings"
	"testing"

	"airline/backend/internal/domain/model"
)

func TestGenerateUniqueSeats_HappyPath(t *testing.T) {
	sg := NewSeatGenerator()

	aircrafts := []model.AircraftType{model.ATR, model.Airbus320, model.Boeing737Max}

	for _, ac := range aircrafts {
		seats, err := sg.GenerateUniqueSeats(ac)
		if err != nil {
			t.Fatalf("unexpected error for %s: %v", ac, err)
		}
		if len(seats) != 3 {
			t.Fatalf("expected 3 seats for %s, got %d", ac, len(seats))
		}

		// Ensure uniqueness
		if seats[0] == seats[1] || seats[0] == seats[2] || seats[1] == seats[2] {
			t.Fatalf("expected unique seats for %s, got %v", ac, seats)
		}

		// Basic format checks: each seat should be like "<number><letter>" and non-empty
		for _, s := range seats {
			if strings.TrimSpace(s) == "" {
				t.Fatalf("empty seat returned for %s", ac)
			}
			// simple validation: must start with a digit
			if s[0] < '0' || s[0] > '9' {
				t.Fatalf("seat %q for %s does not start with a digit", s, ac)
			}
		}
	}
}

func TestGenerateUniqueSeats_InvalidAircraft(t *testing.T) {
	sg := NewSeatGenerator()
	_, err := sg.GenerateUniqueSeats(model.AircraftType("UNKNOWN"))
	if err == nil {
		t.Fatal("expected error for invalid aircraft type, got nil")
	}
	if !strings.Contains(err.Error(), "invalid aircraft type") {
		t.Fatalf("unexpected error message: %v", err)
	}
}
