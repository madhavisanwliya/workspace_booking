package migration

import (
	"context"
	"fmt"
)

// CreateRoleTable ...
func CreateWorkspaceTable() {

	r, err := DbPool.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS workspaces (
		id serial PRIMARY KEY,
		name VARCHAR ( 50 ) NOT NULL,
		floor_id INTEGER REFERENCES floors (id),
		type VARCHAR ( 50 ) NOT NULL,
		active boolean NOT NULL DEFAULT TRUE,
		capacity INT NOT NULL DEFAULT 1,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP)
	`)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

}
