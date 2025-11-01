package service

import (
	"airline/backend/internal/application/dto"
	"airline/backend/internal/domain/model"
	"context"
	"testing"
)

// Mock repository for testing
type mockVoucherRepository struct {
	exists bool
}

// Save implements repository.IVoucherRepository.
func (m *mockVoucherRepository) Save(ctx context.Context, assignment *model.VoucherAssignment) error {
	panic("unimplemented")
}

func (m *mockVoucherRepository) FindByFlight(ctx context.Context, flightNumber string, date string) (bool, error) {
	return m.exists, nil
}

func TestCheck(t *testing.T) {
	repo := &mockVoucherRepository{exists: true}
	seatGen := &SeatGenerator{}
	service := NewVoucherService(repo, seatGen)

	req := dto.CheckRequest{FlightNumber: "FL123", Date: "2025-11-01"}
	res, err := service.Check(context.Background(), req)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !res.Exists {
		t.Fatalf("expected voucher to exist")
	}
}

func TestGenerate(t *testing.T) {
	repo := &mockVoucherRepository{exists: false}
	seatGen := &SeatGenerator{}
	service := NewVoucherService(repo, seatGen)

	req := dto.GenerateRequest{FlightNumber: "FL123", Date: "2025-11-01"}
	_, err := service.Generate(context.Background(), req)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	// Add assertions for the response as needed
}

func TestCheckExistingVoucher(t *testing.T) {
	repo := &mockVoucherRepository{exists: true}
	seatGen := &SeatGenerator{}
	service := NewVoucherService(repo, seatGen)

	req := dto.CheckRequest{FlightNumber: "FL123", Date: "2025-11-01"}
	res, err := service.Check(context.Background(), req)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !res.Exists {
		t.Fatalf("expected voucher to exist")
	}
}

func TestGenerateNewVoucher(t *testing.T) {
	repo := &mockVoucherRepository{exists: false}
	seatGen := &SeatGenerator{}
	service := NewVoucherService(repo, seatGen)

	req := dto.GenerateRequest{FlightNumber: "FL123", Date: "2025-11-01"}
	_, err := service.Generate(context.Background(), req)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	// Add assertions for the response as needed
}
