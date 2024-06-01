package httpHandler

import (
	"authorization-service/internal/interfaces"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

//TODO implement http handler to handle the request from the client and call the service method

type authzHandler struct {
	authzService interfaces.AuthorizationService
}

func NewHTTPAuthzHandler(authzService interfaces.AuthorizationService) *authzHandler {
	return &authzHandler{authzService: authzService}
}

func (h *authzHandler) AssignRole(c echo.Context) error {
	username := c.Param("username")
	// role := c.Param("role")

	err := h.authzService.AssignRole(c.Request().Context(), username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// func (h *authzHandler) CheckPermission(c echo.Context) error {
// 	username := c.Param("username")
// 	permission := c.Param("permission")

// 	allowed, err := h.authzService.CheckPermission(c.Request().Context(), username, permission)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, allowed)
// }

func (h *authzHandler) HandleUserAuthenticatedEvent(ctx context.Context, message string) error {
	// Handle the user authenticated event
	// You can use the echo.Context to access the request and response objects
	// Example:
	// username := ctx.Value("username").(string)
	// role := ctx.Value("role").(string)
	// err := h.authzService.AssignRole(ctx, username, role)
	// if err != nil {
	// 	return err
	// }
	return nil
}
