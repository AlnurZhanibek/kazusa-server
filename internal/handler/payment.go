package handler

import (
	"encoding/base64"
	"encoding/json"
	"github.com/AlnurZhanibek/kazusa-server/internal/service"
	"github.com/google/uuid"
	"log"
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
	OrderID  uuid.UUID `json:"orderId"`
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
		OrderID:  body.OrderID,
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

type PaymentConfirmBody struct {
	Data string `json:"data"`
}
type PaymentData struct {
	OrderID uuid.UUID `json:"order_id"`
	Status  string    `json:"operation_status"`
}
type ConfirmResponse struct {
	OK    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}

func (h *PaymentHandler) Confirm(w http.ResponseWriter, r *http.Request) {
	body := new(PaymentConfirmBody)

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		json.NewEncoder(w).Encode(CreateResponse{
			OK:    false,
			Error: err.Error(),
		})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data, err := base64.StdEncoding.DecodeString(body.Data)
	if err != nil {
		log.Printf("error decoding base64 data %v", err)
		json.NewEncoder(w).Encode(ConfirmResponse{
			OK:    false,
			Error: err.Error(),
		})
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Print(data)

	paymentData := new(PaymentData)
	err = json.Unmarshal(data, paymentData)
	if err != nil {
		json.NewEncoder(w).Encode(ConfirmResponse{
			OK:    false,
			Error: err.Error(),
		})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if paymentData.Status != "success" {
		log.Printf("payment confirmation for %v was not successful: %v", paymentData.OrderID, paymentData.Status)
		json.NewEncoder(w).Encode(ConfirmResponse{
			OK:    false,
			Error: "payment confirmation was not successful",
		})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.service.Confirm(paymentData.OrderID)
	if err != nil {
		json.NewEncoder(w).Encode(ConfirmResponse{
			OK:    false,
			Error: err.Error(),
		})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CreateResponse{
		OK:    true,
		Error: "",
	})
	return
}
