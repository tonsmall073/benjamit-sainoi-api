package main

import (
	"benjamit/src/v1/product"
	"benjamit/src/v1/user" // เปลี่ยนตามชื่อโปรเจกต์ของคุณ
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"benjamit/db"

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
	fmt.Println("Do you want to start server api Press (S)")
	fmt.Println("Do you want to Migrate Press (M)")
	fmt.Print("Please press the button you want to run. : ")

	input, _ := reader.ReadString('\n')
	pressed := strings.TrimSpace(input) // ลบ whitespace และ newline

	if pressed == "S" || pressed == "s" {
		startServer()
	} else if pressed == "M" || pressed == "m" {
		startMigrateDB()
	} else {
		fmt.Println("Invalid input terminates.")
	}
}

func startServer() {
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
