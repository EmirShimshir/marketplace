package test

import (
	"context"
	emailProvider "github.com/EmirShimshir/marketplace/internal/adapter/emailProvider/goMailProvider"
	"github.com/EmirShimshir/marketplace/internal/core/port"
	"github.com/stretchr/testify/require"
	"gopkg.in/gomail.v2"
	"testing"
)

func TestGoMailProvider(t *testing.T) {
	t.Run("test GoMailProvider", func(t *testing.T) {
		senderEmail := "944591@gmail.com"
		password := "vypj wakt oneu fhqs"

		dialer := gomail.NewDialer("smtp.gmail.com", 587, senderEmail, password)
		goMailProvider := emailProvider.NewGoMailProvider(dialer)

		err := goMailProvider.SendEmail(context.Background(), port.CartEmailProviderParam{
			Email:   "emir2701@yandex.ru",
			Subject: "hello",
			Body:    "Word",
		})

		require.Equal(t, true, err == nil)
	})
}
