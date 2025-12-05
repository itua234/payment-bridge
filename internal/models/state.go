package models

import "slices"

//Custom type for payment lifecycle stages
type PaymentState string

const (
	Pending    PaymentState = "pending"
	Authorized PaymentState = "authorized"
	Captured   PaymentState = "captured"
	Voided     PaymentState = "voided"
	Refunded   PaymentState = "refunded"
	Failed     PaymentState = "failed"
)

var validTransitions = map[PaymentState][]PaymentState{
	Pending:    {Authorized},
	Authorized: {Captured, Voided},
	Captured:   {Refunded},
	Voided:     {},
	Refunded:   {},
	Failed:     {},
}

func (p PaymentState) CanTransitionTo(newState PaymentState) bool {
	allowed := validTransitions[p]
	return len(allowed) > 0 && slices.Contains(allowed, newState)
}
