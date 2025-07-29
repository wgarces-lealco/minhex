package handlers

import (
	"encoding/json"
	"net/http"

	"minhex/src/usecases/create_user"
	"minhex/src/usecases/get_user"
)

type UserHandler struct {
	createUserUC create_user.CreateUserUseCase
	getUserUC    get_user.GetUserUseCase
}

func NewUserHandler(createUserUC create_user.CreateUserUseCase, getUserUC get_user.GetUserUseCase) *UserHandler {
	return &UserHandler{
		createUserUC: createUserUC,
		getUserUC:    getUserUC,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req create_user.Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.createUserUC.Execute(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	req := get_user.Request{ID: userID}
	response, err := h.getUserUC.Execute(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
