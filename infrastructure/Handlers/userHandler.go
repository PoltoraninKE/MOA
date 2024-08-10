package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"MOA/services"
	"log/slog"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	userService *services.UserService
	logger      *slog.Logger
}

func NewUserHandler(userService *services.UserService, logger *slog.Logger) *UserHandler {
	return &UserHandler{userService: userService, logger: logger}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	userIdvalue, err := strconv.ParseInt(userID, 10, 64)

	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := h.userService.Read(userIdvalue)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users/{userID}", h.GetUser).Methods("GET")
}
