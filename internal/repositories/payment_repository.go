package repositories

import (
	"github.com/itua234/payment-bridge/internal/models"
	"gorm.io/gorm"
)

type IPaymentRepository interface {
	Create(payment *models.Payment) error
	FindByID(paymentID string) (*models.Payment, error)
	FindByIdempotencyKey(idempotencyKey string) (*models.Payment, error)
	FindByCustomerID(customerID string) ([]models.Payment, error)
	FindByOrderID(orderID string) (*models.Payment, error)
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

func (r *PaymentRepository) FindByCustomerID(customerID string) ([]models.Payment, error) {
	var payments []models.Payment
	result := r.db.Where("customer_id = ?", customerID).Find(&payments)
	return payments, result.Error
}

func (r *PaymentRepository) FindByOrderID(orderID string) (*models.Payment, error) {
	var payment models.Payment
	result := r.db.Where("order_id = ?", orderID).First(&payment)
	return &payment, result.Error
}
