package main

import (
	"bjm/utils"

	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"bjm/db"
	con "bjm/db/benjamit"

	v1 "bjm/src/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
)

// @title Swagger API Docs
// @version 1.0
// @description -
// @BasePath /v1
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
	pressed := strings.TrimSpace(input) // ลบ whitespace และ newline

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

func startServerApi() {
	app := fiber.New()
	useSwagger(app)
	useFiberCors(app)
	v1.UseRoute(app)

	app.Listen(":" + os.Getenv("SERVER_POST"))
}

func useFiberCors(app *fiber.App) {
	config := cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "*",
	}
	app.Use(cors.New(config))
}

func useSwagger(app *fiber.App) {
	app.Get("/v1/swagger/*", swagger.HandlerDefault)

	// เส้นทางสำหรับ swagger.json
	app.Get("/v1/swagger.json", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/v1/swagger.json") // ส่ง swagger.json
	})
}

func startMigrateDB() {
	db.Migrate()
}

func startSeeder() {
	db.Seeder()
}

func startDropAllTablesInDbBenjanit() {
	context, _ := con.Connect()
	utils.DropAllTables(context)
}
