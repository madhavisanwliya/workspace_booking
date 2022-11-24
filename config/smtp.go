package config

import (
	"os"
)

type SmtpConfig struct {
	port     string
	server   string
	password string
	username string
}

var smtpConfig = SmtpConfig{
	port:     "EMAIL_PORT",
	server:   "EMAIL_SERVER",
	username: "EMAIL_USERNAME",
	password: "EMAIL_PASSWORD",
}

func GetEmailPassword() string {
	return os.Getenv(smtpConfig.password)
}

func GetEmailSever() string {
	return os.Getenv(smtpConfig.server)
}

func GetEmailPort() string {
	return os.Getenv(smtpConfig.port)
}

func GetEmailUsername() string {
	return os.Getenv(smtpConfig.username)
}
