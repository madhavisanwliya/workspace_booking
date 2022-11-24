package config

import "time"

// GetEmailDomain returns email's subdomain
func GetEmailDomain() string {
	return "indiumsoft.com"
}

// SqlTimeFormat returns date in YYYY-MM-DD
func SqlTimeFormat(date time.Time) string {
	return date.Format("2006-01-02")
}
