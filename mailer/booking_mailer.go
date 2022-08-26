package mailer

import (
	"fmt"
	"os"
	"workspace_booking/model"
)

func BookingMailer(booking *model.Booking, timing *model.BookingTiming) {

	templatePath := "/text/email-template.html"
	particitpants := model.GetBookingParticipantsDetailsByBookingId(booking.Id)

	recipients := make([]*model.Recipient, 0)

	for _, participant := range particitpants {
		recipient := new(model.Recipient)
		recipient.Name = participant.UserName
		recipient.Email = participant.UserEmail
		recipients = append(recipients, recipient)
	}

	subject := booking.Purpose

	message := "This would informed you that meeting take place on " + timing.FromDate

	const (
		layoutISO = "2006-01-02"
		layoutUS  = "Monday, Jan 2 2006"
	)

	bookingData, _ := model.FetchBooking(booking.Id)

	date := bookingData.FromDateTime

	formatDate := date.Format(layoutUS)

	fmt.Println(formatDate)

	baseUrl := os.Getenv("BASE_URL")

	templateData := map[string]interface{}{
		"Message":           message,
		"Purpose":           booking.Purpose,
		"StartTime":         timing.StartTime,
		"EndTime":           timing.EndTime,
		"Date":              formatDate,
		"City":              bookingData.CityName,
		"Building":          bookingData.BuildingName,
		"Floor":             bookingData.FloorName,
		"WorkspaceName":     bookingData.BookingWorkspace[len(bookingData.BookingWorkspace)-1].WorkspaceName,
		"WorkspaceCapacity": bookingData.BookingWorkspace[len(bookingData.BookingWorkspace)-1].WorkspaceCapacity,
		"BaseUrl":           baseUrl,
	}

	Mailer(recipients, subject, templatePath, message, templateData)

}
