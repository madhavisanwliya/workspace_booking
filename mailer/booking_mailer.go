package mailer

import (
	"os"
	"strings"
	"workspace_booking/config"
	"workspace_booking/model"
)

func BookingMailer(bookingId int16, reminder bool) {

	templatePath := config.GeBookingTemplatePath()
	particitpants := model.GetBookingParticipantsDetailsByBookingId(bookingId)

	recipients := make([]*model.Recipient, 0)
	bookingData, err := model.FetchBooking(bookingId)
	if err != nil {
		panic(err)
	}
	for _, participant := range particitpants {
		recipient := new(model.Recipient)
		recipient.Name = participant.UserName
		recipient.Email = participant.UserEmail
		recipients = append(recipients, recipient)
	}
	if bookingData.CommonEmails != "" {
		commonEmails := strings.SplitAfter(bookingData.CommonEmails, ",")
		for _, commonEmail := range commonEmails {
			commonRecipient := new(model.Recipient)
			commonRecipient.Name = strings.ReplaceAll(commonEmail, " ", "")
			commonRecipient.Email = strings.ReplaceAll(commonEmail, " ", "")
			recipients = append(recipients, commonRecipient)
		}
	}
	layoutUS := config.GetLayoutUS()
	timeLayout := config.GetTimeLayout()

	subject := "Invitation for " + bookingData.Purpose

	if reminder == true {
		subject = "Reminder: " + subject
	}

	date := bookingData.FromDateTime

	formatDate := date.Format(layoutUS)

	message := "This would informed you that meeting take place " + formatDate

	StartTime := date.Format(timeLayout)

	toDate := bookingData.ToDateTime

	EndTime := toDate.Format(timeLayout)

	baseUrl := os.Getenv("BASE_URL")

	templateData := map[string]interface{}{
		"Message":           message,
		"Purpose":           bookingData.Purpose,
		"StartTime":         StartTime,
		"EndTime":           EndTime,
		"Date":              formatDate,
		"City":              bookingData.CityName,
		"Building":          bookingData.BuildingName,
		"Floor":             bookingData.FloorName,
		"WorkspaceName":     bookingData.BookingWorkspace[len(bookingData.BookingWorkspace)-1].WorkspaceName,
		"WorkspaceCapacity": bookingData.BookingWorkspace[len(bookingData.BookingWorkspace)-1].WorkspaceCapacity,
		"BaseUrl":           baseUrl,
		"UserName":          bookingData.UserName,
	}

	Mailer(recipients, subject, templatePath, message, templateData)

}
