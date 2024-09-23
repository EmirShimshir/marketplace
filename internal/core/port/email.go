package port

import "context"

type IEmailService interface {
	Start(ctx context.Context) error
}
