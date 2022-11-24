package migration

import (
	"context"
	"fmt"
)

// CreateRoleTable ...
func CreateAmenityTable() {

	r, err := DbPool.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS amenities (
		id serial PRIMARY KEY,
		name VARCHAR ( 50 ) NOT NULL,
		is_present boolean NOT NULL DEFAULT TRUE,
		workspace_id INTEGER REFERENCES workspaces (id),
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP)
	`)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

}
