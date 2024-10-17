package handler

import (
	"encoding/json"
	"github.com/AlnurZhanibek/kazusa-server/internal/service"
	"github.com/google/uuid"
	"net/http"
)

type PaymentHandler struct {
	service *service.PaymentService
}

func NewPaymentHandler(service *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		service: service,
	}
}

type PaymentCreateBody struct {
	UserID   uuid.UUID `json:"userId"`
	CourseID uuid.UUID `json:"courseId"`
}

type CreateResponse struct {
	OK    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}

func (h *PaymentHandler) Create(w http.ResponseWriter, r *http.Request) {
	body := new(PaymentCreateBody)

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		json.NewEncoder(w).Encode(CreateResponse{
			OK:    false,
			Error: err.Error(),
		})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.service.Create(&service.PaymentCreateBody{
		UserID:   body.UserID,
		CourseID: body.CourseID,
	})
	if err != nil {
		json.NewEncoder(w).Encode(CreateResponse{
			OK:    false,
			Error: err.Error(),
		})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(CreateResponse{
		OK:    true,
		Error: "",
	})
	return
}
