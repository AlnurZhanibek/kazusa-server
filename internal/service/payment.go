package service

import (
	"fmt"
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

type Payment struct {
	ID       uuid.UUID
	UserID   uuid.UUID
	CourseID uuid.UUID
} // @name Payment

type PaymentFilters struct {
	UserID   *uuid.UUID
	CourseID *uuid.UUID
}

func (s *PaymentService) Read(filters PaymentFilters) (*Payment, error) {
	repoPayment, err := s.repo.Read(&repository.PaymentFilters{
		UserID:   filters.UserID,
		CourseID: filters.CourseID,
	})

	if err != nil {
		return nil, fmt.Errorf("service error when reading activity: %v", err)
	}

	payment := &Payment{}
	if repoPayment != nil {
		payment = &Payment{
			ID:       repoPayment.ID,
			UserID:   repoPayment.UserID,
			CourseID: repoPayment.CourseID,
		}
	} else {
		payment = nil
	}

	return payment, nil
}
