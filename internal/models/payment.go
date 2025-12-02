package models

import (
	"time"

	state "gitgub.com/itua234/payment-gateway/models/state"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Payment struct {
	ID             string             `gorm:"type:char(36);primaryKey" json:"id"`
	IdempotencyKey string             `json:"idempotency_key" gorm:"uniqueIndex"`
	Amount         int64              `json:"amount"`
	Currency       string             `gorm:"type:varchar(255)" json:"currency"`
	State          state.PaymentState `json:"state" gorm:"index"`
	BankReference  *string            `json:"bank_reference,omitempty"`
	RetryCount     int                `json:"retry_count"`
	LastError      *string            `json:"last_error,omitempty"`
	CreatedAt      time.Time          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time          `gorm:"autoUpdateTime" json:"updated_at"`
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
