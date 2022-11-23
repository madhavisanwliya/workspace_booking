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

func GetLayoutISO() string {
	return "2006-01-02"
}

func GetLayoutUS() string {
	return "Monday, Jan 2 2006"
}

func GetTimeLayout() string {
	return "15:04 PM"
}

func GetCancelBookingTemplatePath() string {
	return "/text/cancelbooking-template.html"
}

func GeBookingTemplatePath() string {
	return "/text/email-template.html"
}
