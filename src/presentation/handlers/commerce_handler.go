package handlers

import (
	"encoding/json"
	"net/http"

	"minhex/src/usecases/activate_commerce"
	"minhex/src/usecases/create_commerce"
)

type CommerceHandler struct {
	createCommerceUC   create_commerce.CreateCommerceUseCase
	activateCommerceUC activate_commerce.ActivateCommerceUseCase
}

func NewCommerceHandler(createCommerceUC create_commerce.CreateCommerceUseCase, activateCommerceUC activate_commerce.ActivateCommerceUseCase) *CommerceHandler {
	return &CommerceHandler{
		createCommerceUC:   createCommerceUC,
		activateCommerceUC: activateCommerceUC,
	}
}

func (h *CommerceHandler) CreateCommerce(w http.ResponseWriter, r *http.Request) {
	var req create_commerce.Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.createCommerceUC.Execute(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *CommerceHandler) ActivateCommerce(w http.ResponseWriter, r *http.Request) {
	var req activate_commerce.Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.activateCommerceUC.Execute(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
