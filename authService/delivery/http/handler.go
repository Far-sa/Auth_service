package httpServer

import (
	"authentication-service/domain/param"
	"authentication-service/interfaces"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	authService interfaces.AuthenticationService
}

func NewHTTPAuthHandler(authService interfaces.AuthenticationService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var userReq param.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.authService.Login(r.Context(), userReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
