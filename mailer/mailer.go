package mailer

// gmail service
import (
	"fmt"
	"html/template"
	"io"
	"os"
	"strconv"
	"workspace_booking/config"
	"workspace_booking/model"

	gomail "gopkg.in/mail.v2"
)

func Mailer(to []*model.Recipient, subject, templatePath, message string, templateData interface{}) {
	from := config.GetEmailUsername()
	password := config.GetEmailPassword()
	smtpHost := config.GetEmailSever()
	smtpPort, _ := strconv.Atoi(config.GetEmailPort())

	m := gomail.NewMessage()

	addresses := make([]string, len(to))
	for i, user := range to {
		addresses[i] = m.FormatAddress(user.Email, user.Name)
	}

	m.SetHeader("From", config.GetEmailUsername())
	m.SetHeader("To", addresses...)
	m.SetHeader("Subject", subject)

	pwd, _ := os.Getwd()

	// m.Attach(pwd + templatePath)

	t, _ := template.ParseFiles(pwd + templatePath)
	m.AddAlternativeWriter("text/html", func(w io.Writer) error {
		return t.Execute(w, templateData)
	})

	d := gomail.NewDialer(smtpHost, smtpPort, from, password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	fmt.Println("Email Sent Successfully!")

}

func SendOtpMail(to []*model.Recipient, subject, message string) {
	from := config.GetEmailUsername()
	password := config.GetEmailPassword()
	smtpHost := config.GetEmailSever()
	smtpPort, _ := strconv.Atoi(config.GetEmailPort())

	m := gomail.NewMessage()

	addresses := make([]string, len(to))
	for i, user := range to {
		addresses[i] = m.FormatAddress(user.Email, user.Name)
	}

	m.SetHeader("From", config.GetEmailUsername())
	m.SetHeader("To", addresses...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", message)

	d := gomail.NewDialer(smtpHost, smtpPort, from, password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	fmt.Println("Verification mail Sent Successfully!")

}
