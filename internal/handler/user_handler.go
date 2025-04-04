package handler

import (
	"libra-ry/internal/domain"
	"libra-ry/internal/usecase"
	"libra-ry/pkg"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserHandler struct {
	useCase usecase.UserUseCase
}

func NewUserHandler(useCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{useCase: useCase}
}

// GetAllUsers godoc
// @Summary Get all users
// @Tags users
// @Produce json
// @Param page query int false "Page number"
// @Param username query string false "Filter by username"
// @Param role query string false "Filter by role"
// @Param permissions query string false "Filter by permission"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /user [get]
func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	username := c.Query("username")
	role := c.Query("role")
	permissionsParam := c.Query("permissions") // misal: "buku_read,user_read"
	var permissions []string
	if permissionsParam != "" {
		for _, p := range strings.Split(permissionsParam, ",") {
			p = strings.TrimSpace(p)
			if p != "" {
				permissions = append(permissions, p)
			}
		}
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 100
	offset := (page - 1) * limit

	users, total, err := h.useCase.GetAll(username, role, permissions, limit, offset)
	if err != nil {
		return pkg.ErrorResponse(c, 500, err.Error())
	}

	responseUsers := pkg.MapUsersToResponse(users)
	return pkg.SuccessResponse(c, "Success", responseUsers, total)
}

// GetUserByID godoc
// @Summary Get user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Security BearerAuth
// @Success 200 {object} domain.UserResponse
// @Router /user/{id} [get]
func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return pkg.ErrorResponse(c, 400, "Invalid ID")
	}
	user, err := h.useCase.GetByID(uint(id))
	if err != nil {
		return pkg.ErrorResponse(c, 404, "User not found")
	}

	responseUser := pkg.MapUserToResponse(*user)
	return pkg.SuccessResponse(c, "Success", responseUser, 1)
}

// CreateUser godoc
// @Summary Create new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body domain.UserSwagger true "User data"
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Router /user [post]
func (h *UserHandler) Create(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return pkg.ErrorResponse(c, 400, "Invalid request body")
	}
	if err := h.useCase.Create(&user); err != nil {
		return pkg.ErrorResponse(c, 500, err.Error())
	}
	return pkg.SuccessResponse(c, "Success", fiber.Map{"message": "User created successfully"}, 1)
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

// ChangePassword godoc
// @Summary Change user password
// @Tags users
// @Accept json
// @Produce json
// @Param body body ChangePasswordRequest true "Change Password"
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 400,401,500 {object} map[string]string
// @Router /user/change-password [put]
func (h *UserHandler) ChangePassword(c *fiber.Ctx) error {
	var req ChangePasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return pkg.ErrorResponse(c, fiber.StatusBadRequest, "invalid request")
	}

	userToken := c.Locals("user").(jwt.MapClaims)
	userID := userToken["user_id"]
	// parse userID to uint
	userIDUint, ok := userID.(float64)
	if !ok {
		return pkg.ErrorResponse(c, fiber.StatusUnauthorized, "invalid user id")
	}
	userIDInt := uint(userIDUint)
	err := h.useCase.ChangePassword(userIDInt, req.OldPassword, req.NewPassword)
	if err != nil {
		return pkg.ErrorResponse(c, fiber.StatusUnauthorized, err.Error())
	}

	return pkg.SuccessResponse(c, "password updated successfully", nil, 1)
}

// UpdateUser godoc
// @Summary Update user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body domain.UserResponse true "Updated user data"
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Router /user/{id} [put]
func (h *UserHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return pkg.ErrorResponse(c, 400, "Invalid ID")
	}
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return pkg.ErrorResponse(c, 400, "Invalid request body")
	}
	user.ID = uint(id)
	if err := h.useCase.Update(&user); err != nil {
		return pkg.ErrorResponse(c, 500, err.Error())
	}
	return pkg.SuccessResponse(c, "Success", fiber.Map{"message": "User updated successfully"}, 1)
}

// DeleteUser godoc
// @Summary Delete user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Router /user/{id} [delete]
func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return pkg.ErrorResponse(c, 400, "Invalid ID")
	}
	if err := h.useCase.Delete(uint(id)); err != nil {
		return pkg.ErrorResponse(c, 500, err.Error())
	}
	return pkg.SuccessResponse(c, "Success", fiber.Map{"message": "User deleted successfully"}, 1)
}
