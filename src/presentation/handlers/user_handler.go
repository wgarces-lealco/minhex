package handlers

import (
	"encoding/json"
	"minhex/src/usecases/create_user"
	"net/http"
)

type UserHandler struct {
	createUserUC *create_user.UseCase
}

func NewUserHandler(createUserUC *create_user.UseCase) *UserHandler {
	return &UserHandler{
		createUserUC: createUserUC,
	}
}

type CreateUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type CreateUserResponse struct {
	UserID string `json:"user_id"`
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.createUserUC.Execute(r.Context(), create_user.Request{
		Email: req.Email,
		Name:  req.Name,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CreateUserResponse{
		UserID: response.UserID,
	})
}
