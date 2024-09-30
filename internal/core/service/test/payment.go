package test

import (
	"github.com/EmirShimshir/marketplace/internal/core/domain"
)

// PaymentPayloadBuilder helps to build PaymentPayload objects with default values.
type PaymentPayloadBuilder struct {
	payload domain.PaymentPayload
}

// NewPaymentPayloadBuilder returns a new instance of PaymentPayloadBuilder.
func NewPaymentPayloadBuilder() *PaymentPayloadBuilder {
	return &PaymentPayloadBuilder{
		payload: domain.PaymentPayload{
			OrderID: domain.NewID(),
			PaySum:  1000,
		},
	}
}

// WithOrderID sets the OrderID field.
func (b *PaymentPayloadBuilder) WithOrderID(orderID domain.ID) *PaymentPayloadBuilder {
	b.payload.OrderID = orderID
	return b
}

// WithPaySum sets the PaySum field.
func (b *PaymentPayloadBuilder) WithPaySum(paySum int64) *PaymentPayloadBuilder {
	b.payload.PaySum = paySum
	return b
}

// Build returns the built PaymentPayload.
func (b *PaymentPayloadBuilder) Build() domain.PaymentPayload {
	return b.payload
}
