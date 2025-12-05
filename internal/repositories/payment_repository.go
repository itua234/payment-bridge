package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/itua234/payment-bridge/internal/models"
	"gorm.io/gorm"
)

var (
	ErrPaymentNotFound  = errors.New("payment not found")
	ErrDuplicatePayment = errors.New("duplicate payment")
)

type IPaymentRepository interface {
	Create(ctx context.Context, payment *models.Payment) error
	FindByID(ctx context.Context, paymentID string) (*models.Payment, error)
	FindByIdempotencyKey(ctx context.Context, key string) (*models.Payment, error)
	FindByCustomerID(ctx context.Context, customerID string) ([]models.Payment, error)
	FindByOrderID(ctx context.Context, orderID string) (*models.Payment, error)
	IncrementRetryCount(ctx context.Context, paymentID string) error
}

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) IPaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) Create(ctx context.Context, payment *models.Payment) error {
	return r.db.WithContext(ctx).Create(payment).Error
}

func (r *PaymentRepository) FindByID(ctx context.Context, paymentID string) (*models.Payment, error) {
	var payment models.Payment
	result := r.db.WithContext(ctx).First(&payment, "id = ?", paymentID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrPaymentNotFound
		}
		return nil, fmt.Errorf("find payment: %w", result.Error)
	}
	return &payment, result.Error
}

func (r *PaymentRepository) FindByIdempotencyKey(ctx context.Context, idempotencyKey string) (*models.Payment, error) {
	var payment models.Payment
	result := r.db.WithContext(ctx).First(&payment, "idempotency_key = ?", idempotencyKey)
	return &payment, result.Error
}

func (r *PaymentRepository) FindByCustomerID(ctx context.Context, customerID string) ([]models.Payment, error) {
	var payments []models.Payment
	result := r.db.WithContext(ctx).Where("customer_id = ?", customerID).Find(&payments)
	return payments, result.Error
}

func (r *PaymentRepository) FindByOrderID(ctx context.Context, orderID string) (*models.Payment, error) {
	var payment models.Payment
	result := r.db.WithContext(ctx).First(&payment, "order_id = ?", orderID)
	return &payment, result.Error
}

func (r *PaymentRepository) IncrementRetryCount(
	ctx context.Context,
	paymentID string,
) error {
	result := r.db.WithContext(ctx).
		Model(&models.Payment{}).
		Where("id = ?", paymentID).
		UpdateColumn("retry_count", gorm.Expr("retry_count + 1"))

	if result.Error != nil {
		return fmt.Errorf("increment retry count: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return ErrPaymentNotFound
	}

	return nil
}
