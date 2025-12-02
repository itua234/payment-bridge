package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StateTransition struct {
	ID        string       `gorm:"type:char(36);primaryKey" json:"id"`
	PaymentID string       `gorm:"type:char(36);index;not null" json:"payment_id"`
	From      PaymentState `gorm:"type:varchar(50);not null" json:"from"`
	To        PaymentState `gorm:"type:varchar(50);not null;index" json:"to"`
	Timestamp time.Time    `gorm:"autoCreateTime;not null;index" json:"timestamp"`

	// Foreign key relationship
	//Payment Payment `gorm:"foreignKey:PaymentID;constraint:OnDelete:CASCADE" json:"-"`
}

func (StateTransition) TableName() string {
	return "state_transitions"
}

func (s *StateTransition) BeforeCreate(tx *gorm.DB) error {
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	return nil
}
