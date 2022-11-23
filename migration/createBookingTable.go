package migration

import (
	"context"
	"fmt"
)

// CreateBookingsTable ...
func CreateBookingsTable() {

	r, err := DbPool.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS bookings (
    id serial PRIMARY KEY,
		city_id int,
		building_id int,
		floor_id int,
		user_id int,
		from_datetime TIMESTAMP NOT NULL,
		to_datetime TIMESTAMP NOT NULL,
		comments text,
		common_emails text,
		active boolean NOT NULL DEFAULT TRUE,
		purpose VARCHAR ( 255 ),
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP)
	`)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

}
