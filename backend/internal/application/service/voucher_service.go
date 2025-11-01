package service

import (
	"airline/backend/internal/application/dto"
	"airline/backend/internal/domain/model"
	"airline/backend/internal/domain/repository"
	"context"
	"fmt"
)

// Custom error
var ErrVoucherAlreadyExists = fmt.Errorf("vouchers have already been generated for this flight and date")

// VoucherService handles the application's use cases.
type VoucherService struct {
	repo    repository.IVoucherRepository
	seatGen *SeatGenerator
}

// NewVoucherService creates a new VoucherService.
func NewVoucherService(repo repository.IVoucherRepository, seatGen *SeatGenerator) *VoucherService {
	return &VoucherService{
		repo:    repo,
		seatGen: seatGen,
	}
}

// Check handles the voucher existence check.
func (s *VoucherService) Check(ctx context.Context, req dto.CheckRequest) (*dto.CheckResponse, error) {
	exists, err := s.repo.FindByFlight(ctx, req.FlightNumber, req.Date)
	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}
	return &dto.CheckResponse{Exists: exists}, nil
}

// Generate handles the new voucher generation.
func (s *VoucherService) Generate(ctx context.Context, req dto.GenerateRequest) (*dto.GenerateResponse, error) {
	// 1. Check if vouchers already exist (critical check)
	exists, err := s.repo.FindByFlight(ctx, req.FlightNumber, req.Date)
	if err != nil {
		return nil, fmt.Errorf("database check failed: %w", err)
	}
	if exists {
		return nil, ErrVoucherAlreadyExists
	}

	// 2. Convert string aircraft type to domain model
	aircraftType := model.AircraftType(req.Aircraft)

	// 3. Generate 3 unique seats
	seats, err := s.seatGen.GenerateUniqueSeats(aircraftType)
	if err != nil {
		return nil, fmt.Errorf("seat generation failed: %w", err)
	}

	// 4. Create the domain entity
	assignment := model.NewVoucherAssignment(
		req.Name,
		req.ID,
		req.FlightNumber,
		req.Date,
		aircraftType,
		seats,
	)

	// 5. Save to repository
	if err := s.repo.Save(ctx, assignment); err != nil {
		return nil, fmt.Errorf("database save failed: %w", err)
	}

	// 6. Return successful response
	return &dto.GenerateResponse{
		Success: true,
		Seats:   seats,
	}, nil
}
