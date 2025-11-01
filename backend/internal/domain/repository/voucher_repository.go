package repository

import (
	"airline/backend/internal/domain/model"
	"context"
)

// IVoucherRepository defines the persistence interface for voucher assignments.
type IVoucherRepository interface {
	// FindByFlight checks if an assignment exists for a given flight number and date.
	FindByFlight(ctx context.Context, flightNumber, date string) (bool, error)

	// Save persists a new voucher assignment.
	Save(ctx context.Context, assignment *model.VoucherAssignment) error
}
