package model

import (
	"context"
	"fmt"
	"time"
	"workspace_booking/migration"
)

// BookingParticipant struct
type BookingParticipant struct {
	Id           int16     `json:"id"`
	BookingId    int16     `json:"booking_id"`
	UserId       int16     `json:"user_id"`
	FloorId      int16     `json:"floor_id"`
	FromDateTime string    `json:"from_datetime"`
	ToDateTime   string    `json:"to_datetime"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
type BookingParticipantDetail struct {
	Id        int16  `json:"id"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
}
type BookingParticipantDetails struct {
	BookingParticipantDetails []*BookingParticipantDetail
}

func (bp *BookingParticipant) CreateBookingParticipant() error {
	dt := time.Now()
	query := "INSERT INTO booking_participants (booking_id, user_id, created_at, updated_at, floor_id, from_datetime, to_datetime) VALUES ($1, $2, $3, $4, $5, $6, $7) " +
		"RETURNING id, created_at, updated_at"
	d := migration.DbPool.QueryRow(context.Background(), query, &bp.BookingId, &bp.UserId, dt, dt, &bp.FloorId, &bp.FromDateTime, &bp.ToDateTime)
	err := d.Scan(&bp.Id, &bp.CreatedAt, &bp.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func BulkInsertBookingParticipant(booking *Booking) error {
	for _, userId := range booking.UserIds {
		bookingParticipant := new(BookingParticipant)
		bookingParticipant.BookingId = booking.Id
		bookingParticipant.FloorId = booking.FloorId
		bookingParticipant.FromDateTime = booking.FromDateTime
		bookingParticipant.ToDateTime = booking.ToDateTime
		bookingParticipant.UserId = userId
		err := bookingParticipant.CreateBookingParticipant()
		if err != nil {
			return err
		}
	}
	return nil
}

func BulkDeleteBookingParticipant(bookingId int16) error {
	participants, e := migration.DbPool.Query(context.Background(), "SELECT * from booking_participants where booking_id = $1", bookingId)

	defer participants.Close()

	for participants.Next() {
		participant := new(BookingParticipantDetail)
		e = participants.Scan(&participant.Id, &participant.UserName, &participant.UserEmail)
		DeleteBookingParticipant(participant.Id)
	}
	if e != nil {
		return e
	}
	return nil
}

func DeleteBookingParticipant(Id int16) error {
	ok, err := migration.DbPool.Exec(context.Background(), "DELETE FROM booking_participants WHERE id=$1", Id)
	fmt.Println(ok)
	return err
}

func GetBookingParticipantsDetailsByBookingId(bookingId int16) []*BookingParticipantDetail {
	// query all booking_participants data
	participants, e := migration.DbPool.Query(context.Background(), "SELECT user_id, (select name from users where id = booking_participants.user_id) as user_name, (select email from users where id = booking_participants.user_id) as user_email from booking_participants where booking_id = $1", bookingId)

	defer participants.Close()
	// declare BookingParticipantDetail array variable
	bookingParticipantsDetails := make([]*BookingParticipantDetail, 0)

	// iterate over booking_participants
	for participants.Next() {
		participant := new(BookingParticipantDetail)
		e = participants.Scan(&participant.Id, &participant.UserName, &participant.UserEmail)
		bookingParticipantsDetails = append(bookingParticipantsDetails, participant)
	}
	if e != nil {
		fmt.Println("Failed to get bookings_details record :", e)
		return []*BookingParticipantDetail{}
	}
	return bookingParticipantsDetails
}
