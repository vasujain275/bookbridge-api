package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vasujain275/bookbridge-api/internal/repository"
	"github.com/vasujain275/bookbridge-api/internal/service"
	"github.com/vasujain275/bookbridge-api/internal/util"
)

// UserHandler handles HTTP requests for users.
type UserHandler struct {
	service service.UserService
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

// CreateUserRequest represents the expected request payload for creating a user.
type CreateUserRequest struct {
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	Role      string `json:"role" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

// UpdateUserRequest represents the expected request payload for updating a user.
type UpdateUserRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

// GetUser godoc
// @Summary Get user by ID
// @Description Get a user by its ID.
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} util.Response "User found"
// @Failure 400 {object} util.Response "Invalid ID supplied"
// @Failure 404 {object} util.Response "User not found"
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		util.SendBadRequest(c, "Invalid user ID", err.Error())
		return
	}

	user, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		util.SendNotFound(c, err.Error())
		return
	}
	util.SendOK(c, "User found", user)
}

// ListUsers godoc
// @Summary List users
// @Description Get a paginated list of users.
// @Tags users
// @Accept json
// @Produce json
// @Param limit query int false "Limit" default(10)
// @Param offset query int false "Offset" default(0)
// @Success 200 {object} util.Response "List of users"
// @Failure 500 {object} util.Response "Internal server error"
// @Router /users [get]
func (h *UserHandler) ListUsers(c *gin.Context) {
	// Default values
	limit := 10
	offset := 0

	if l := c.Query("limit"); l != "" {
		if lim, err := strconv.Atoi(l); err == nil {
			limit = lim
		}
	}
	if o := c.Query("offset"); o != "" {
		if off, err := strconv.Atoi(o); err == nil {
			offset = off
		}
	}

	users, err := h.service.List(c.Request.Context(), int32(limit), int32(offset))
	if err != nil {
		util.SendInternalServerError(c, err.Error())
		return
	}
	util.SendOK(c, "Users retrieved", users)
}

// CreateUser godoc
// @Summary Create user
// @Description Create a new user.
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User data"
// @Success 201 {object} util.Response "User created successfully"
// @Failure 400 {object} util.Response "Invalid request"
// @Failure 500 {object} util.Response "Internal server error"
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.SendBadRequest(c, "Invalid request body", err.Error())
		return
	}

	// Build params for the repository layer.
	params := repository.CreateUserParams{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: req.Password,
		Role:         req.Role,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
	}

	user, err := h.service.Create(c.Request.Context(), params)
	if err != nil {
		util.SendInternalServerError(c, err.Error())
		return
	}
	util.SendCreated(c, "User created", user)
}

// UpdateUser godoc
// @Summary Update user
// @Description Update an existing user.
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body UpdateUserRequest true "Updated user data"
// @Success 200 {object} util.Response "User updated successfully"
// @Failure 400 {object} util.Response "Invalid request"
// @Failure 404 {object} util.Response "User not found"
// @Failure 500 {object} util.Response "Internal server error"
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		util.SendBadRequest(c, "Invalid user ID", err.Error())
		return
	}

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.SendBadRequest(c, "Invalid request body", err.Error())
		return
	}

	params := repository.UpdateUserParams{
		ID:           id,
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: req.Password,
	}

	user, err := h.service.Update(c.Request.Context(), params)
	if err != nil {
		util.SendInternalServerError(c, err.Error())
		return
	}
	util.SendOK(c, "User updated", user)
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete a user by its ID.
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 204 {object} util.Response "User deleted successfully"
// @Failure 400 {object} util.Response "Invalid user ID"
// @Failure 500 {object} util.Response "Internal server error"
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		util.SendBadRequest(c, "Invalid user ID", err.Error())
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		util.SendInternalServerError(c, err.Error())
		return
	}
	util.SendNoContent(c)
}
