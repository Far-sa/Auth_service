package httpHandler

import (
	"encoding/json"
	"net/http"
	"user-service/internal/interfaces"
	"user-service/internal/param"
)

type UserHandler struct {
	userService interfaces.UserService
}

func NewHTTPAuthHandler(userService interfaces.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	var userReq param.GetUserByEmail
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.userService.GetUserByEmail(r.Context(), userReq.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
