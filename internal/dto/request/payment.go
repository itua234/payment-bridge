package request

type CreatePaymentRequest struct {
	IdempotencyKey string `json:"idempotency_key" binding:"required"`
	Amount         int64  `json:"amount" binding:"required,gt=0"`
	Currency       string `json:"currency" binding:"required,len=3"`
}
