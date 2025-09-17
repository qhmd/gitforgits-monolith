package http

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/qhmd/gitforgits/config"
	authDto "github.com/qhmd/gitforgits/internal/dto/auth"
	"github.com/qhmd/gitforgits/internal/middleware"
	"github.com/qhmd/gitforgits/internal/usecase"
	"github.com/qhmd/gitforgits/utils"
)

type UserHandler struct {
	uc *usecase.UsersUseCase
}

func NewHandlerUser(app *fiber.App, uc *usecase.UsersUseCase) {
	h := &UserHandler{uc: uc}
	app.Get("/admin/users/:id<^[0-9]+$>", middleware.JWT(), middleware.IsAdmin(), h.GetUserByID)
	app.Get("/admin/users/", middleware.JWT(), middleware.IsAdmin(), h.GetListUsers)
	app.Put("/admin/users/:id<^[0-9]+$>", middleware.JWT(), middleware.IsAdmin(), middleware.ValidateUserUpdate(), h.UpdateUsers)
	app.Delete("admin/users/:id<^[0-9]+$>", middleware.JWT(), middleware.IsAdmin(), h.DeleteUserByID)
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Retrieve a single user by its ID
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Security ApiKeyAuth
// @Param id path int true "Users ID"
// @Success 400 {object} users.InvalidId
// @Success 200 {object} users.SuccessGetUser
// @Failure 404 {object} users.UserNotFoundResponse
// @Router /admin/users/{id} [get]
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "invalid id", nil)
	}
	data, err := h.uc.GetUserByID(c.Context(), id)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "user not found", nil)
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "successfully get the user", data)
}

// GetListUsers godoc
// @Summary Get List User
// @Description Get list user
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Security ApiKeyAuth
// @Success 200 {array} users.SuccessGetList
// @Failure 404 {object} users.UserNotFoundResponse
// @Router /admin/users/ [get]
func (h *UserHandler) GetListUsers(c *fiber.Ctx) error {
	data, err := h.uc.ListUser(c.Context())
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "something went wrong", nil)
	}
	return utils.SuccessResponse(c, fiber.StatusOK, "successfully get list user", data)
}

// UpdateUsers godoc
// @Summary Update user
// @Description Update user
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Security ApiKeyAuth
// @Param id path int true "User ID"
// @Param Auth body users.UpdateRequest true "Update Account"
// @Success 200 {object} users.SuccessGetList
// @Success 400 {object} users.InvalidId
// @Failure 404 {object} users.UserNotFoundResponse
// @Failure 409 {object} users.EmailAlreadyUsed
// @Failure 500 {object} users.ErrorResponse
// @Router /admin/users/{id} [put]
func (h *UserHandler) UpdateUsers(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "invalid id", nil)
	}
	user := c.Locals("validateUser").(authDto.UserResponse)

	_, err = h.uc.GetUserByID(c.Context(), id)
	fmt.Print(err)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "user not found", nil)

	}

	updateUser, err := h.uc.UpdateUser(c.Context(), &user, id)
	if err != nil {
		if errors.Is(err, config.ErrUserExists) {
			return utils.ErrorResponse(c, fiber.StatusConflict, "try another email", "email already used")
		}
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "something went wrong", nil)
	}
	fmt.Print("updated user : ", updateUser)
	return utils.SuccessResponse(c, fiber.StatusOK, "successfully update the user", updateUser)
}

// DeleteUserByID godoc
// @Summary Delete user
// @Description Delete user by id
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Security ApiKeyAuth
// @Param id path int true "User ID"
// @Success 200 {object} users.SuccessDeleteUser
// @Success 400 {object} users.InvalidId
// @Failure 404 {object} users.UserNotFoundResponse
// @Failure 500 {object} users.ErrorResponse
// @Router /admin/users/{id} [delete]
func (h *UserHandler) DeleteUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "invalid id", nil)
	}
	_, err = h.uc.GetUserByID(c.Context(), id)
	fmt.Print(err)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "user not found", nil)

	}
	if err = h.uc.DeleteUser(c.Context(), id); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "something went wrong", nil)

	}
	return utils.SuccessResponse(c, fiber.StatusOK, "success delete user", nil)
}
