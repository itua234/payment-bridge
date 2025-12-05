package bank

import "time"

type AuthorizeRequest struct {
	Amount      int64  `json:"amount"`
	CardNumber  string `json:"card_number"`
	CVV         string `json:"cvv"`
	ExpiryMonth int64  `json:"expiry_month"`
	ExpiryYear  int64  `json:"expiry_year"`
}

type AuthorizeResponse struct {
	Amount          int64     `json:"amount"`
	AuthorizationID string    `json:"authorization_id"`
	CreatedAt       time.Time `json:"created_at"`
	Currency        string    `json:"currency"`
	ExpiresAt       time.Time `json:"expires_at"`
	Status          string    `json:"status"`
}

type CaptureRequest struct {
	Amount          int64  `json:"amount"`
	AuthorizationID string `json:"authorization_id"`
}

type CaptureResponse struct {
	Amount          int64     `json:"amount"`
	AuthorizationID string    `json:"authorization_id"`
	CaptureID       string    `json:"capture_id"`
	CapturedAt      time.Time `json:"captured_at"`
	Currency        string    `json:"currency"`
	Status          string    `json:"status"`
}

type VoidRequest struct {
	AuthorizationID string `json:"authorization_id"`
}

type VoidResponse struct {
	AuthorizationID string    `json:"authorization_id"`
	Status          string    `json:"status"`
	VoidID          string    `json:"void_id"`
	VoidedAt        time.Time `json:"voided_at"`
}

type RefundRequest struct {
	Amount    int64  `json:"amount"`
	CaptureID string `json:"capture_id"`
}

type RefundResponse struct {
	Amount     int64     `json:"amount"`
	CaptureID  string    `json:"capture_id"`
	Currency   string    `json:"currency"`
	RefundID   string    `json:"refund_id"`
	RefundedAt time.Time `json:"refunded_at"`
	Status     string    `json:"status"`
}
