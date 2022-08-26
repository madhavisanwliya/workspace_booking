package main

import (
	"context"
	"fmt"
	"log"
	"workspace_booking/config"
	db "workspace_booking/migration"
	"workspace_booking/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	dbPool := db.GetDbConnectionPool()
	// close database
	defer dbPool.Close()
	// check db
	err := dbPool.Ping(context.Background())
	CheckError(err)

	fmt.Println("Connected!")

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	app.Static("/static", "./assets")

	router.SetupRoutes(app)

	println(config.GetServerPort())
	log.Fatalln(app.Listen(config.GetServerPort()))
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
