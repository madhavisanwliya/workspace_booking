package mailer

import (
	"os"
	"workspace_booking/model"
)

func BookingMailer(bookingId int16) {

	templatePath := "/text/email-template.html"
	particitpants := model.GetBookingParticipantsDetailsByBookingId(bookingId)

	recipients := make([]*model.Recipient, 0)

	for _, participant := range particitpants {
		recipient := new(model.Recipient)
		recipient.Name = participant.UserName
		recipient.Email = participant.UserEmail
		recipients = append(recipients, recipient)
	}

	const (
		layoutISO  = "2006-01-02"
		layoutUS   = "Monday, Jan 2 2006"
		timeLayout = "15:04 PM"
	)

	bookingData, _ := model.FetchBooking(bookingId)

	subject := "Invitation for " + bookingData.Purpose

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
