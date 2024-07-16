package handler

import (
	"encoding/json"
	"kazusa-server/internal/service"
	"net/http"
	"time"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	OK    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	var credentials LoginRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		encoder.Encode(LoginResponse{
			OK:    false,
			Error: "",
		})
	}

	token, err := h.service.Login(credentials.Email, credentials.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		encoder.Encode(LoginResponse{
			OK:    false,
			Error: "service error",
		})
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		Secure:   true,
		HttpOnly: true,
	})

	encoder.Encode(LoginResponse{
		OK:    true,
		Error: "",
	})
}

type RegisterRequest struct {
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Phone                string `json:"phone"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

type RegisterResponse struct {
	OK    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	var credentials RegisterRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&credentials)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		encoder.Encode(RegisterResponse{
			OK:    false,
			Error: "decode error",
		})
	}

	token, err := h.service.Register(credentials.Name, credentials.Email, credentials.Phone, credentials.Password, credentials.PasswordConfirmation)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		encoder.Encode(RegisterResponse{
			OK:    false,
			Error: "token gen error",
		})
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		Secure:   true,
		HttpOnly: true,
	})

	encoder.Encode(RegisterResponse{
		OK:    true,
		Error: "",
	})
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: "",
	})
}
