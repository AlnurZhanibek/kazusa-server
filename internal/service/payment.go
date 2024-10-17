package service

import (
	"github.com/AlnurZhanibek/kazusa-server/internal/repository"
	"github.com/google/uuid"
)

type PaymentService struct {
	repo *repository.PaymentRepository
}

func NewPaymentService(repo *repository.PaymentRepository) *PaymentService {
	return &PaymentService{repo: repo}
}

type PaymentCreateBody struct {
	UserID   uuid.UUID
	CourseID uuid.UUID
}

func (s *PaymentService) Create(payment *PaymentCreateBody) error {
	return s.repo.Create(&repository.PaymentCreateBody{
		UserID:   payment.UserID,
		CourseID: payment.CourseID,
	})
}
