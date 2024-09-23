package port

import "context"

type CartEmailProviderParam struct {
	Email  string
	Subject string
	Body    string
}

type IEmailProvider interface {
	SendEmail(ctx context.Context, param CartEmailProviderParam) error
}
