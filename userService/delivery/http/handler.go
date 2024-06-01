package httpHandler

import (
	"net/http"
	"user-service/internal/interfaces"
	"user-service/internal/param"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService interfaces.UserService
}

func NewHTTPAuthHandler(userService interfaces.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) SignUp(c echo.Context) error {
	var userReq param.RegisterRequest
	if err := c.Bind(&userReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, err := h.userService.Register(c.Request().Context(), userReq)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, resp)

}

func (h *UserHandler) GetUserByEmail(c echo.Context) error {
	var userReq param.GetUserByEmail
	if err := c.Bind(&userReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, err := h.userService.GetUserByEmail(c.Request().Context(), userReq.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
