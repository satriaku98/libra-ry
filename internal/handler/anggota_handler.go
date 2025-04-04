package handler

import (
	"errors"
	"libra-ry/internal/domain"
	"libra-ry/internal/usecase"
	"libra-ry/pkg"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type AnggotaHandler struct {
	usecase usecase.AnggotaUseCase
}

func NewAnggotaHandler(u usecase.AnggotaUseCase) *AnggotaHandler {
	return &AnggotaHandler{usecase: u}
}

// GetAnggota godoc
// @Summary Get member detail
// @Tags anggota
// @Security BearerAuth
// @Produce json
// @Success 200 {object} domain.AnggotaSwagger
// @Router /anggota [get]
func (h *AnggotaHandler) Get(c *fiber.Ctx) error {
	// get user id from JWT token
	userToken := c.Locals("user").(jwt.MapClaims)
	userID := userToken["user_id"]
	// parse userID to uint
	userIDUint, ok := userID.(float64)
	if !ok {
		return pkg.ErrorResponse(c, fiber.StatusUnauthorized, "invalid user id")
	}
	userIDInt := uint(userIDUint)

	anggota, err := h.usecase.GetByUserID(userIDInt)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return pkg.ErrorResponse(c, fiber.StatusNotFound, "anggota not found")
		}
		return pkg.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.SuccessResponse(c, http.StatusText(fiber.StatusOK), anggota, 1)
}

// CreateAnggota godoc
// @Summary Create member profile
// @Tags anggota
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body domain.AnggotaSwagger true "Anggota data"
// @Success 201 {object} domain.AnggotaSwagger
// @Router /anggota [post]
func (h *AnggotaHandler) Create(c *fiber.Ctx) error {
	// get user id from JWT token
	userToken := c.Locals("user").(jwt.MapClaims)
	userID := userToken["user_id"]
	// parse userID to uint
	userIDUint, ok := userID.(float64)
	if !ok {
		return pkg.ErrorResponse(c, fiber.StatusUnauthorized, "invalid user id")
	}
	userIDInt := uint(userIDUint)

	var anggota domain.Anggota
	if err := c.BodyParser(&anggota); err != nil {
		return pkg.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	anggota.UserID = userIDInt
	err := h.usecase.Create(&anggota)
	if err != nil {
		return pkg.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.SuccessResponse(c, http.StatusText(fiber.StatusCreated), anggota, 1)
}

// UpdateAnggota godoc
// @Summary Update member profile
// @Tags anggota
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body domain.AnggotaSwagger true "Anggota data"
// @Success 200 {object} domain.AnggotaSwagger
// @Router /anggota [put]
func (h *AnggotaHandler) Update(c *fiber.Ctx) error {
	// get user id from JWT token
	userToken := c.Locals("user").(jwt.MapClaims)
	userID := userToken["user_id"]
	// parse userID to uint
	userIDUint, ok := userID.(float64)
	if !ok {
		return pkg.ErrorResponse(c, fiber.StatusUnauthorized, "invalid user id")
	}
	userIDInt := uint(userIDUint)

	var anggota domain.Anggota
	if err := c.BodyParser(&anggota); err != nil {
		return pkg.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	anggota.UserID = userIDInt
	err := h.usecase.Update(&anggota)
	if err != nil {
		return pkg.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.SuccessResponse(c, http.StatusText(fiber.StatusOK), anggota, 1)
}

// DeleteAnggota godoc
// @Summary Delete member profile
// @Tags anggota
// @Security BearerAuth
// @Produce json
// @Success 200 {object} string
// @Router /anggota [delete]
func (h *AnggotaHandler) Delete(c *fiber.Ctx) error {
	// get user id from JWT token
	userToken := c.Locals("user").(jwt.MapClaims)
	userID := userToken["user_id"]
	// parse userID to uint
	userIDUint, ok := userID.(float64)
	if !ok {
		return pkg.ErrorResponse(c, fiber.StatusUnauthorized, "invalid user id")
	}
	userIDInt := uint(userIDUint)

	err := h.usecase.DeleteByUserID(userIDInt)
	if err != nil {
		return pkg.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return pkg.SuccessResponse(c, http.StatusText(fiber.StatusAccepted), "Anggota deleted successfully", 1)
}
