package services

import (
	"context"
	"errors"

	"github.com/itua234/payment-bridge/internal/bank"
	request "github.com/itua234/payment-bridge/internal/dto/request"
	models "github.com/itua234/payment-bridge/internal/models"
	"github.com/itua234/payment-bridge/internal/repositories"
	"gorm.io/gorm"
)

type IPaymentService interface {
	CreatePayment(ctx context.Context, req request.CreatePaymentRequest) (*models.Payment, error)
}

type PaymentService struct {
	bankClient  *bank.Client
	paymentRepo repositories.IPaymentRepository
}

func NewPaymentService(
	bankClient *bank.Client,
	paymentRepo repositories.IPaymentRepository,
) *PaymentService {
	return &PaymentService{
		bankClient:  bankClient,
		paymentRepo: paymentRepo,
	}
}

func (s *PaymentService) CreatePayment(
	ctx context.Context,
	req request.CreatePaymentRequest,
) (*models.Payment, error) {
	existing, err := s.paymentRepo.FindByIdempotencyKey(ctx, req.IdempotencyKey)
	if err == nil {
		return existing, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	payment := &models.Payment{
		IdempotencyKey: req.IdempotencyKey,
		Amount:         req.Amount,
		Currency:       req.Currency,
		State:          models.Pending,
		//MetaData:       metadataInstance,
	}

	if err := s.paymentRepo.Create(ctx, payment); err != nil {
		return nil, err
	}

	return payment, nil
}
