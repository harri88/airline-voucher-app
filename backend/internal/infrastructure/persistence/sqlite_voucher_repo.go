package persistence

import (
	"airline/backend/internal/domain/model"
	"context"
	"database/sql"
	"time"
)

// SQLiteVoucherRepository is the concrete implementation of IVoucherRepository.
type SQLiteVoucherRepository struct {
	db *sql.DB
}

// NewSQLiteVoucherRepository creates a new repository and returns the concrete type.
// Returning the concrete type (instead of the interface) lets Wire use wire.Bind
// in the provider set without producing duplicate bindings.
func NewSQLiteVoucherRepository(db *sql.DB) *SQLiteVoucherRepository {
	return &SQLiteVoucherRepository{db: db}
}

// FindByFlight uses parameterized queries to check for existence.
func (r *SQLiteVoucherRepository) FindByFlight(ctx context.Context, flightNumber, date string) (bool, error) {
	var count int
	query := "SELECT COUNT(1) FROM voucher_assignments WHERE flight_number = ? AND flight_date = ?"

	err := r.db.QueryRowContext(ctx, query, flightNumber, date).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// Save uses parameterized queries to insert a new record.
func (r *SQLiteVoucherRepository) Save(ctx context.Context, a *model.VoucherAssignment) error {
	query := `
		INSERT INTO voucher_assignments 
		(crew_name, crew_id, flight_number, flight_date, aircraft_type, seat1, seat2, seat3, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query,
		a.CrewName,
		a.CrewID,
		a.FlightNumber,
		a.FlightDate,
		a.AircraftType,
		a.Seat1,
		a.Seat2,
		a.Seat3,
		a.CreatedAt.Format(time.RFC3339),
	)
	return err
}
