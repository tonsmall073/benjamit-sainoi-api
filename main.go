package main

import (
	v1 "bjm/src/http/v1"
	v2 "bjm/src/http/v2"
	"bjm/utils"
	"net"

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
	"google.golang.org/grpc"

	gRpc "bjm/src/grpc"

	_ "bjm/docs" // swagger docs
	"bjm/middlewares"
)

func main() {
	mode := os.Getenv("APP_MODE")
	fmt.Println("App Mode : " + mode)
	if mode == "production" {
		if err := godotenv.Load("./envs/.env.production"); err != nil {
			log.Fatal("[ERROR] loading .env.production file")
		}
	} else if mode == "staging" {
		if err := godotenv.Load("./envs/.env.staging"); err != nil {
			log.Fatal("[ERROR] loading .env.staging file")
		}
	} else {
		if err := godotenv.Load("./envs/.env.development"); err != nil {
			log.Fatal("[ERROR] loading .env.development file")
		}
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Do you want to Start Http and gRpc Server API Press (HG)")
	fmt.Println("Do you want to Start Http Server API Press (H)")
	fmt.Println("Do you want to Start gRpc Server API Press (G)")
	fmt.Println("Do you want to Migrate Press (M)")
	fmt.Println("Do you want to Database Seeding Press (I)")
	fmt.Println("Do you want to Drop ALL Tables in database name Benjamit Press (D)")
	fmt.Print("Please press the key you want to run. : ")

	input, _ := reader.ReadString('\n')
	pressed := strings.TrimSpace(input)
	if pressed == "HG" || pressed == "hg" {
		startServerApi()
		startGrpcServer()
	} else if pressed == "H" || pressed == "h" {
		startServerApi()
	} else if pressed == "G" || pressed == "g" {
		startGrpcServer()
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
	middlewares.UseFiberHelmet(app)
	middlewares.UseRequestLimit(app)
	middlewares.UseSwagger(app)
	sse := ssefiber.New(app, "")
	route := middlewares.UseApiTransactionLog(app)
	utils.UseValidator()
	v1.UseRoute(route, sse)
	v2.UseRoute(route)

	if err := app.Listen(":" + os.Getenv("HTTP_SERVER_PORT")); err != nil {
		log.Fatalf("failed to start http server: %v", err)
	}
}

func startGrpcServer() {
	lis, err := net.Listen("tcp", ":"+os.Getenv("GRPC_SERVER_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middlewares.UseApiTransactionLogGrpc(),
			middlewares.UseTimeZoneGrpc(),
		),
	)

	gRpc.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start gRpc server: %v", err)
	}
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
