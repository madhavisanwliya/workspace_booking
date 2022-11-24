package model

import (
	"context"
	"fmt"
	"time"
	"workspace_booking/migration"
)

// BookingWorkspace struct
type BookingWorkspace struct {
	Id           int16     `json:"id"`
	BookingId    int16     `json:"booking_id"`
	WorkspaceId  int16     `json:"workspace_id"`
	FloorId      int16     `json:"floor_id"`
	FromDateTime string    `json:"from_datetime"`
	ToDateTime   string    `json:"to_datetime"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type BookingWorkspaceDetail struct {
	Id                int16  `json:"id"`
	WorkspaceName     string `json:"workspace_name"`
	WorkspaceType     string `json:"workspace_type"`
	WorkspaceCapacity int    `json:"workspace_capacity"`
}

type BookingWorkspaceDetails struct {
	BookingWorkspaceDetails []*BookingWorkspaceDetail
}

func (bw *BookingWorkspace) CreateBookingWorkspace() error {
	dt := time.Now()
	query := "INSERT INTO booking_workspaces (booking_id, workspace_id, created_at, updated_at, floor_id, from_datetime, to_datetime) VALUES ($1, $2, $3, $4, $5, $6, $7) " +
		"RETURNING id, created_at, updated_at"
	d := migration.DbPool.QueryRow(context.Background(), query, &bw.BookingId, &bw.WorkspaceId, dt, dt, &bw.FloorId, &bw.FromDateTime, &bw.ToDateTime)
	err := d.Scan(&bw.Id, &bw.CreatedAt, &bw.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func BulkInsertBookingWorkspace(booking *Booking, timing *BookingTiming) error {
	for _, selectWorkspace := range booking.SelectedWorkspaces {
		timing.FromDate = selectWorkspace.Date
		timing.ToDate = selectWorkspace.Date
		fromDatetTime, toDateTime := BookingTimestamp(timing)
		for _, seatId := range selectWorkspace.Seats {
			bookingWorkspace := new(BookingWorkspace)
			bookingWorkspace.BookingId = booking.Id
			bookingWorkspace.FloorId = booking.FloorId
			bookingWorkspace.FromDateTime = fromDatetTime
			bookingWorkspace.ToDateTime = toDateTime
			bookingWorkspace.WorkspaceId = seatId
			err := bookingWorkspace.CreateBookingWorkspace()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func GetBookingWorkspacesDetailsByBookingId(bookingId int16) []*BookingWorkspaceDetail {
	workspaces, e := migration.DbPool.Query(context.Background(), "SELECT workspace_id, (select name from workspaces where id = booking_workspaces.workspace_id) as workspace_name, (select type from workspaces where id = booking_workspaces.workspace_id) as workspace_type, (select capacity from workspaces where id = booking_workspaces.workspace_id) as workspace_capacity from booking_workspaces where booking_id = $1", bookingId)

	defer workspaces.Close()
	bookingWorkspaceDetails := make([]*BookingWorkspaceDetail, 0)

	for workspaces.Next() {
		workspace := new(BookingWorkspaceDetail)
		e = workspaces.Scan(&workspace.Id, &workspace.WorkspaceName, &workspace.WorkspaceType, &workspace.WorkspaceCapacity)
		bookingWorkspaceDetails = append(bookingWorkspaceDetails, workspace)
	}

	if e != nil {
		fmt.Println("Failed to get bookings_details record :", e)
		return []*BookingWorkspaceDetail{}
	}
	return bookingWorkspaceDetails
}
