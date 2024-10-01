package main

import (
	"bjm/src/v1/product"
	"bjm/src/v1/user"

	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"bjm/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
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
	fmt.Print("Please press the key you want to run. : ")

	input, _ := reader.ReadString('\n')
	pressed := strings.TrimSpace(input) // ลบ whitespace และ newline

	if pressed == "S" || pressed == "s" {
		startServerApi()
	} else if pressed == "M" || pressed == "m" {
		startMigrateDB()
	} else if pressed == "I" || pressed == "i" {
		StartSeeding()
	} else {
		fmt.Println("Invalid input terminates.")
	}
}

func startServerApi() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "*",
	}))

	route := app.Group("/v1")
	user.Setup(route)
	product.Setup(route)

	app.Listen(":" + os.Getenv("SERVER_POST"))
}

func startMigrateDB() {
	db.Migrate()
}

func StartSeeding() {
	db.Seed()
}
