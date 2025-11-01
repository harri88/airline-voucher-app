package dto

// CheckRequest is the DTO for checking voucher existence.
type CheckRequest struct {
	FlightNumber string `json:"flightNumber" binding:"required"`
	Date         string `json:"date" binding:"required"` // Expects YYYY-MM-DD
}

// CheckResponse is the DTO for the check result.
type CheckResponse struct {
	Exists bool `json:"exists"`
}

// GenerateRequest is the DTO for generating new vouchers.
type GenerateRequest struct {
	Name         string `json:"name" binding:"required"`
	ID           string `json:"id" binding:"required"`
	FlightNumber string `json:"flightNumber" binding:"required"`
	Date         string `json:"date" binding:"required"`     // Expects YYYY-MM-DD
	Aircraft     string `json:"aircraft" binding:"required"` // "ATR", "Airbus 320", etc.
}

// GenerateResponse is the DTO for a successful generation.
type GenerateResponse struct {
	Success bool     `json:"success"`
	Seats   []string `json:"seats"`
}

// ErrorResponse is a generic DTO for error messages.
type ErrorResponse struct {
	Error string `json:"error"`
}
