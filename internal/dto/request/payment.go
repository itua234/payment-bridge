package request

type CreatePaymentRequest struct {
	IdempotencyKey string `json:"idempotency_key" binding:"required"`
	Amount         int64  `json:"amount" binding:"required,gt=0"`
	Currency       string `json:"currency" binding:"required,len=3"`
}

type AuthorizeRequest struct {
	CardNumber  string `json:"card_number"`
	CVV         string `json:"cvv"`
	ExpiryMonth string `json:"expiry_month"`
	ExpiryYear  string `json:"expiry_year"`
	Amount      int64  `json:"amount"`
	OrderID     string `json:"order_id"`
	CustomerID  string `json:"customer_id"`
}

type CaptureRequest struct {
	Amount          int64  `json:"amount"`
	AuthorizationID string `json:"authorization_id"`
}
