package migration

import (
	"context"
	"fmt"
)

// CreateBookingWorkspacesTable ...
func CreateBookingWorkspaceTable() {

	r, err := DbPool.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS booking_workspaces (
    id serial PRIMARY KEY,
		booking_id int,
		workspace_id INTEGER REFERENCES workspaces (id),
		from_datetime TIMESTAMP NOT NULL,
		to_datetime TIMESTAMP NOT NULL,
		floor_id int,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP)
	`)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

}
