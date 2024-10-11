package main

import (
	v1 "bjm/src/v1"
	v2 "bjm/src/v2"
	"bjm/utils"

	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"bjm/db"
	con "bjm/db/benjamit"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/jsorb84/ssefiber"

	_ "bjm/docs"
	"bjm/middlewares"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Do you want to Start Server API Press (S)")
	fmt.Println("Do you want to Migrate Press (M)")
	fmt.Println("Do you want to Database Seeding Press (I)")
	fmt.Println("Do you want to Drop ALL Tables in database name Benjamit Press (D)")
	fmt.Print("Please press the key you want to run. : ")

	input, _ := reader.ReadString('\n')
	pressed := strings.TrimSpace(input)

	if pressed == "S" || pressed == "s" {
		startServerApi()
	} else if pressed == "M" || pressed == "m" {
		startMigrateDB()
	} else if pressed == "I" || pressed == "i" {
		startSeeder()
	} else if pressed == "D" || pressed == "d" {
		startDropAllTablesInDbBenjanit()
	} else {
		fmt.Println("Invalid input terminates.")
	}
}

// @title Swagger API Docs
// @version 1.0
// @description Benjamit API

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" <JWT token>.
func startServerApi() {
	app := fiber.New()
	middlewares.UseTimeZone(app)
	middlewares.UseFiberCors(app)
	middlewares.UseSwagger(app)
	sse := ssefiber.New(app, "")
	route := middlewares.UseApiTransactionLog(app)
	v1.UseRoute(route, sse)
	v2.UseRoute(route)

	app.Listen(":" + os.Getenv("SERVER_POST"))
}

func startMigrateDB() {
	db.Migrate()
}

func startSeeder() {
	db.Seeder()
}

func startDropAllTablesInDbBenjanit() {
	context, contextErr := con.Connect()
	if contextErr != nil {
		log.Fatalf("failed to connect to database: %v", contextErr)
	}
	utils.DropAllTables(context)
}
