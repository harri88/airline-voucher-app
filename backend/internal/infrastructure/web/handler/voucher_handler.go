package handler

import (
	"airline/backend/internal/application/dto"
	"airline/backend/internal/application/service"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// VoucherHandler handles HTTP requests for vouchers.
type VoucherHandler struct {
	service *service.VoucherService
}

// NewVoucherHandler creates a new handler.
func NewVoucherHandler(s *service.VoucherService) *VoucherHandler {
	return &VoucherHandler{service: s}
}

// CheckAssignment
// @Summary Check if vouchers exist for a flight
// @Description Checks if vouchers have already been assigned to a specific flight on a specific date.
// @Tags Vouchers
// @Accept json
// @Produce json
// @Param request body dto.CheckRequest true "Flight Check Request"
// @Success 200 {object} dto.CheckResponse
// @Failure 400 {object} dto.ErrorResponse "Bad Request"
// @Failure 500 {object} dto.ErrorResponse "Internal Server Error"
// @Router /api/check [post]
func (h *VoucherHandler) CheckAssignment(c *gin.Context) {
	var req dto.CheckRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request body"})
		return
	}

	res, err := h.service.Check(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Could not check vouchers"})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GenerateAssignment
// @Summary Generate 3 new voucher seats
// @Description Generates 3 unique random voucher seats for a flight, fails if they already exist.
// @Tags Vouchers
// @Accept json
// @Produce json
// @Param request body dto.GenerateRequest true "Voucher Generation Request"
// @Success 201 {object} dto.GenerateResponse
// @Failure 400 {object} dto.ErrorResponse "Bad Request or Invalid Aircraft"
// @Failure 409 {object} dto.ErrorResponse "Vouchers already generated"
// @Failure 500 {object} dto.ErrorResponse "Internal Server Error"
// @Router /api/generate [post]
func (h *VoucherHandler) GenerateAssignment(c *gin.Context) {
	var req dto.GenerateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid request body: " + err.Error()})
		return
	}

	res, err := h.service.Generate(c.Request.Context(), req)
	if err != nil {
		// Handle specific known error
		if errors.Is(err, service.ErrVoucherAlreadyExists) {
			c.JSON(http.StatusConflict, dto.ErrorResponse{Error: err.Error()})
			return
		}
		// Handle seat generation error (e.g., bad aircraft type)
		if strings.Contains(err.Error(), "seat generation failed") {
			c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
			return
		}
		// Generic internal error
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Could not generate vouchers"})
		return
	}

	c.JSON(http.StatusCreated, res)
}
