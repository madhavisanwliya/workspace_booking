package model

import (
	"context"
	"fmt"
	"time"
	"workspace_booking/migration"
)

// Building struct
type Amenity struct {
	Id            int       `json:"id"`
	Name          string    `json:"name"`
	IsPresent     bool      `json:"is_present"`
	WorkSpaceId   int       `json:"workspace_id"`
	WorksapceName string    `json:"workspace_name"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Amenities struct
type Amenities struct {
	Amenities []*Amenity
}

func GetAllAmenities() []*Amenity {
	rows, err := migration.DbPool.Query(context.Background(), "select amenities.id, amenities.name, amenities.is_present, amenities.workspace_id as workspace_id, workspaces.name as workspace_name, amenities.created_at, amenities.updated_at from amenities LEFT JOIN workspaces ON amenities.workspace_id = workspaces.id")
	defer rows.Close()
	if err != nil {
		fmt.Println("Failed to get amenities record :", err)
		return []*Amenity{}
	}

	// declare empty post variable
	amenities := make([]*Amenity, 0)
	// iterate over rows
	for rows.Next() {
		amenity := new(Amenity)
		err := rows.Scan(&amenity.Id, &amenity.Name, &amenity.IsPresent, &amenity.WorkSpaceId, &amenity.WorksapceName, &amenity.CreatedAt, &amenity.UpdatedAt)
		if err != nil {
			fmt.Println("Failed to get amenities record :", err)
			return []*Amenity{}
		}
		amenities = append(amenities, amenity)
	}
	return amenities
}

func BulkInsertAmenities(workSpaceId int, amenities []Amenity) []Amenity {
	amenitiesArr := make([]Amenity, 0)
	for _, amenity := range amenities {
		amenity.WorkSpaceId = workSpaceId
		err := amenity.CreateAmenity()
		amenitiesArr = append(amenitiesArr, amenity)
		if err != nil {
			return []Amenity{}
		}
	}
	return amenitiesArr
}

func (amenity *Amenity) CreateAmenity() error {
	dt := time.Now()
	query := "INSERT INTO amenities (name, workspace_id, is_present, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, (select name from workspaces where workspaces.id = amenities.workspace_id) as workspace_name, created_at, updated_at"
	data := migration.DbPool.QueryRow(context.Background(), query, &amenity.Name, &amenity.WorkSpaceId, &amenity.IsPresent, dt, dt)
	err := data.Scan(&amenity.Id, &amenity.WorksapceName, &amenity.CreatedAt, &amenity.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func GetAmenityByID(amenityId int) Amenity {
	amenity := Amenity{}
	rows := migration.DbPool.QueryRow(context.Background(), "select amenities.id, amenities.name, amenities.is_present, amenities.workspace_id as workspace_id, workspaces.name as workspace_name, amenities.created_at, amenities.updated_at from amenities LEFT JOIN workspaces ON amenities.workspace_id = workspaces.id where amenities.id = $1", amenityId)
	err := rows.Scan(&amenity.Id, &amenity.Name, &amenity.IsPresent, &amenity.WorkSpaceId, &amenity.WorksapceName, &amenity.CreatedAt, &amenity.UpdatedAt)
	if err != nil {
		fmt.Println("Failed to get amenity record :", err)
		return Amenity{}
	}
	return amenity
}

func GetAmenitiesByWorkSpaceID(WorkSpaceId int) []Amenity {
	rows, err := migration.DbPool.Query(context.Background(), "select amenities.id, amenities.name, amenities.is_present, amenities.workspace_id as workspace_id, workspaces.name as workspace_name, amenities.created_at, amenities.updated_at from amenities LEFT JOIN workspaces ON amenities.workspace_id = workspaces.id where amenities.workspace_id = $1", WorkSpaceId)
	defer rows.Close()
	// declare amenities array
	if err != nil {
		fmt.Println("Failed to get amenities record :", err)
		return []Amenity{}
	}
	amenities := make([]Amenity, 0)
	// iterate over rows
	for rows.Next() {
		amenity := Amenity{}
		err := rows.Scan(&amenity.Id, &amenity.Name, &amenity.IsPresent, &amenity.WorkSpaceId, &amenity.WorksapceName, &amenity.CreatedAt, &amenity.UpdatedAt)
		if err != nil {
			fmt.Println("Failed to get amenities record :", err)
			return []Amenity{}
		}
		amenities = append(amenities, amenity)
	}
	return amenities
}
