package repositories

import (
	"github.com/itua234/payment-gateway/internal/models"
	"gorm.io/gorm"
)

type IPaymentRepository interface {
	Create(payment *models.Payment) error
	FindByID(paymentID string) (*models.Payment, error)
	FindByIdempotencyKey(idempotencyKey string) (*models.Payment, error)
}

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) IPaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) Create(payment *models.Payment) error {
	return r.db.Create(payment).Error
}

func (r *PaymentRepository) FindByID(paymentID string) (*models.Payment, error) {
	var payment models.Payment
	result := r.db.First(&payment, "id = ?", paymentID)
	return &payment, result.Error
}

func (r *PaymentRepository) FindByIdempotencyKey(idempotencyKey string) (*models.Payment, error) {
	var payment models.Payment
	result := r.db.First(&payment, "idempotency_key = ?", idempotencyKey)
	return &payment, result.Error
}
