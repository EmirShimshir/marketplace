package gomail

import (
	"gopkg.in/gomail.v2"
)

type Config struct {
	Host        string
	Port        int
	SenderEmail string
	Password    string
}

func NewClient(cfg *Config) *gomail.Dialer {
	return gomail.NewDialer(cfg.Host, cfg.Port, cfg.SenderEmail, cfg.Password)
}
