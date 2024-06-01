package httpHandler

import (
	"authentication-service/domain/param"
	"authentication-service/interfaces"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService interfaces.AuthenticationService
}

func NewHTTPAuthHandler(authService interfaces.AuthenticationService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c echo.Context) error {
	var userReq param.LoginRequest

	if err := c.Bind(&userReq); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resp, err := h.authService.Login(c.Request().Context(), userReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
