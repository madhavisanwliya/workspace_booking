package model

import (
	"context"
	"fmt"
	"time"
	"workspace_booking/migration"
)

// Role struct
type FloorWorkSpace struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	FloorId   int       `json:"floor_id"`
	Type      string    `json:"type"`
	Active    bool      `json:"active"`
	Capacity  int       `json:"capacity"`
	Amenities []Amenity `json:"amenities"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Roles struct
type FloorWorkSpaces struct {
	FloorWorkSpaces []FloorWorkSpace `json:"workspaces"`
}

// GetAllworkspaces will fetch all the roles from roles table
func GetAllworkspaces(FloorId int) []FloorWorkSpace {
	rows, err := migration.DbPool.Query(context.Background(), "SELECT * FROM workspaces where floor_id = $1", FloorId)
	if err != nil {
		return []FloorWorkSpace{}
	}

	defer rows.Close()

	floorWorkspaces := make([]FloorWorkSpace, 0)
	for rows.Next() {
		floorWorkspace := FloorWorkSpace{}
		err = rows.Scan(&floorWorkspace.Id, &floorWorkspace.Name, &floorWorkspace.FloorId, &floorWorkspace.Type, &floorWorkspace.Active, &floorWorkspace.Capacity, &floorWorkspace.CreatedAt, &floorWorkspace.UpdatedAt)
		amenities := GetAmenitiesByWorkSpaceID(floorWorkspace.Id)
		floorWorkspace.Amenities = amenities
		if err != nil {
			fmt.Println("Failed", err)
			return []FloorWorkSpace{}
		}

		floorWorkspaces = append(floorWorkspaces, floorWorkspace)
	}

	return floorWorkspaces
}

func BulkFloorWorkspacesCreate(FloorId int, floorWorkspaces []FloorWorkSpace) []FloorWorkSpace {
	floorWorkspaceRecords := make([]FloorWorkSpace, 0)
	for _, floorWorkspaceRecord := range floorWorkspaces {
		floorWorkspaceRecord.FloorId = FloorId
		err := floorWorkspaceRecord.CreateFloorWorkspace()
		if err != nil {
			return []FloorWorkSpace{}
		}
		allAmenities := BulkInsertAmenities(floorWorkspaceRecord.Id, floorWorkspaceRecord.Amenities)
		floorWorkspaceRecord.Amenities = allAmenities
		floorWorkspaceRecords = append(floorWorkspaceRecords, floorWorkspaceRecord)
	}
	return floorWorkspaceRecords
}

// CreateFloorWorkspace will create the workspace record in db
func (floorWorkspace *FloorWorkSpace) CreateFloorWorkspace() error {
	dt := time.Now()
	query := "INSERT INTO workspaces (name, floor_id, type, active, capacity, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *"
	d := migration.DbPool.QueryRow(context.Background(), query, &floorWorkspace.Name, &floorWorkspace.FloorId, &floorWorkspace.Type, &floorWorkspace.Active, &floorWorkspace.Capacity, dt, dt)
	err := d.Scan(&floorWorkspace.Id, &floorWorkspace.Name, &floorWorkspace.FloorId, &floorWorkspace.Type, &floorWorkspace.Active, &floorWorkspace.Capacity, &floorWorkspace.CreatedAt, &floorWorkspace.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
