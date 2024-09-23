package emailProvider

import (
	"context"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	"gopkg.in/gomail.v2"
)

type GoMailProvider struct {
	dialer *gomail.Dialer
}

func NewGoMailProvider(dialer *gomail.Dialer) *GoMailProvider {
	return &GoMailProvider{
		dialer: dialer,
	}
}

func (g *GoMailProvider) SendEmail(ctx context.Context, param port.CartEmailProviderParam) error {
	m := gomail.NewMessage()
	m.SetHeader("From", g.dialer.Username)
	m.SetHeader("To", param.Email)
	m.SetHeader("Subject", param.Subject)
	m.SetBody("text/plain", param.Body)

	return g.dialer.DialAndSend(m)
}
