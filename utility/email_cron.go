package utility

import (
	"fmt"
	"time"
	"workspace_booking/mailer"
	"workspace_booking/model"

	"github.com/go-co-op/gocron"
)

func BookingMailCron() {
	s := gocron.NewScheduler(time.UTC)

	s.Every("15m").Do(func() {
		bookingIds := model.GetBookingForReminder()
		fmt.Println(bookingIds)
		for _, id := range bookingIds {
			mailer.BookingMailer(id)
		}

	})

	s.StartBlocking()
}
