package mailer

import (
	"context"
	"math/rand"
	"strconv"
	"time"
	"workspace_booking/migration"
	"workspace_booking/model"
)

func UserMailer(user *model.User) {
	recipient := new(model.Recipient)
	recipient.Name = user.Name
	recipient.Email = user.Email
	recipients := make([]*model.Recipient, 0)
	recipients = append(recipients, recipient)

	otp := rand.Intn(99999)
	subject := "User Verification"
	otpValue := strconv.Itoa(otp)
	message := "Your verification OTP code is <b>" + otpValue + "</b>"
	dt := time.Now()
	_, err := migration.DbPool.Exec(context.Background(), "UPDATE users SET otp=$1, updated_at=$2 WHERE id=$3", otpValue, dt, user.ID)
	if err == nil {
		SendOtpMail(recipients, subject, message)
	}
}
