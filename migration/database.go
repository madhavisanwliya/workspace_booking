package migration

import (
	"context"
	"workspace_booking/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DbPool *pgxpool.Pool

func GetDbConnectionPool() *pgxpool.Pool {
	if DbPool != nil {
		return DbPool
	}

	psqlconn := config.GetDBConnectionURL()
	println(psqlconn)
	db, err := pgxpool.Connect(context.Background(), psqlconn)

	// open database
	checkError(err)
	DbPool = db

	// Don't change the order here
	CreateRoleTable()
	CreateUserTable()
	CreateCityTable()
	// CreateLocationTable()
	CreateBuildingTable()
	CreateFloorTable()
	CreateWorkspaceTable()
	CreateAmenityTable()
	CreateBookingsTable()
	CreateBookingParticipantsTable()
	CreateBookingWorkspaceTable()

	return DbPool
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
