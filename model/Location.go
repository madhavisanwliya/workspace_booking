package model

import (
	"context"
	"fmt"
	"time"
	"workspace_booking/migration"
)

// Building struct
type Location struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	CityId    int64     `json:"city_id"`
	CityName  string    `json:"city_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Buildings struct
type Locations struct {
	Locations []*Location
}

func GetAllLocations() []*Location {
	rows, _ := migration.DbPool.Query(context.Background(), "select locations.id, locations.name, locations.city_id, cities.name as city_name, locations.created_at, locations.updated_at from locations LEFT JOIN cities ON locations.city_id = cities.id")
	defer rows.Close()

	// declare empty post variable
	locations := make([]*Location, 0)
	// iterate over rows
	for rows.Next() {
		location := new(Location)
		err := rows.Scan(&location.Id, &location.Name, &location.CityId, &location.CityName, &location.CreatedAt, &location.UpdatedAt)
		if err != nil {
			fmt.Println("Failed to get locations record :", err)
			return []*Location{}
		}
		locations = append(locations, location)
	}
	return locations
}

func (location *Location) CreateLocation() error {
	dt := time.Now()
	query := "INSERT INTO locations (name, city_id, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at"
	city := migration.DbPool.QueryRow(context.Background(), "select name from cities where id = $1", location.CityId)
	err := city.Scan(&location.CityName)
	if err != nil {
		return err
	}
	d := migration.DbPool.QueryRow(context.Background(), query, &location.Name, &location.CityId, dt, dt)
	err = d.Scan(&location.Id, &location.CreatedAt, &location.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func GetLocationByID(locationId int) Location {
	location := Location{}
	rows := migration.DbPool.QueryRow(context.Background(), "select locations.id, locations.name, locations.city_id, cities.name as city_name, locations.created_at, locations.updated_at from locations LEFT JOIN cities ON locations.city_id = cities.id where locations.id = $1", locationId)
	err := rows.Scan(&location.Id, &location.Name, &location.CityId, &location.CityName, &location.CreatedAt, &location.UpdatedAt)
	if err != nil {
		fmt.Println("Failed to get locations record :", err)
		return Location{}
	}
	return location
}
