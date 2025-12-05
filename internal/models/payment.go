package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	ID             string `gorm:"type:char(36);primaryKey" json:"id"`
	IdempotencyKey string `gorm:"uniqueIndex" json:"idempotency_key"`

	CardNumber  string `json:"card_number"`
	CVV         string `json:"cvv"`
	ExpiryMonth int64  `json:"expiry_month"`
	ExpiryYear  int64  `json:"expiry_year"`

	Amount     int64  `json:"amount"`
	Currency   string `gorm:"type:varchar(255)" json:"currency"`
	CustomerID string `gorm:"index;type:varchar(255)" json:"customer_id"`
	OrderID    string `gorm:"index;type:varchar(255)" json:"order_id"`

	State PaymentState `gorm:"index" json:"state" `

	AuthorizationRef *string `json:"authorization_ref,omitempty"`
	CaptureRef       *string `json:"capture_ref,omitempty"`
	VoidRef          *string `json:"void_ref,omitempty"`
	RefundRef        *string `json:"refund_ref,omitempty"`

	RetryCount int     `json:"retry_count"`
	LastError  *string `json:"last_error,omitempty"`

	//History *[]StateTransition `gorm:"foreignKey:PaymentID" json:"history,omitempty"`

	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	AuthorizedAt *time.Time `json:"authorized_at,omitempty"`
	CapturedAt   *time.Time `json:"captured_at,omitempty"`
	VoidedAt     *time.Time `json:"voided_at,omitempty"`
	RefundedAt   *time.Time `json:"refunded_at,omitempty"`
}

func (Payment) TableName() string {
	return "payments"
}

func (p *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	return nil
}
