package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/ZiplEix/pixel-espion/database"
	"github.com/ZiplEix/pixel-espion/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"

	_ "github.com/ZiplEix/pixel-espion/docs" // Swagger docs
)

func checkEnv() error {
	log.Printf("Checking environment variables...")

	// application
	if _, ok := os.LookupEnv("PORT"); !ok {
		return errors.New("env var 'PORT' is not set")
	}
	if _, ok := os.LookupEnv("VERSION"); !ok {
		return errors.New("env var 'VERSION' is not set")
	}
	// database
	if _, ok := os.LookupEnv("POSTGRES_HOST"); !ok {
		return errors.New("env var 'POSTGRES_HOST' is not set")
	}
	if _, ok := os.LookupEnv("POSTGRES_PORT"); !ok {
		return errors.New("env var 'POSTGRES_PORT' is not set")
	}
	if _, ok := os.LookupEnv("POSTGRES_USER"); !ok {
		return errors.New("env var 'POSTGRES_USER' is not set")
	}
	if _, ok := os.LookupEnv("POSTGRES_PASSWORD"); !ok {
		return errors.New("env var 'POSTGRES_PASSWORD' is not set")
	}
	if _, ok := os.LookupEnv("POSTGRES_DB"); !ok {
		return errors.New("env var 'POSTGRES_DB' is not set")
	}
	// jwt
	if _, ok := os.LookupEnv("JWT_SECRET"); !ok {
		return errors.New("env var 'JWT_SECRET' is not set")
	}

	return nil
}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	err = checkEnv()
	if err != nil {
		panic(err)
	}

	err = database.Connect()
	if err != nil {
		panic(err)
	}

	err = database.Migrate()
	if err != nil {
		panic(err)
	}
}

// @title pixe espion API
// @version 0.1
// @description This is the API for a spy pixel service
// @host localhost:8080
// @BasePath /
// @contact.name ZiplEix
// @contact.email OnGithub
func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("ALLOWED_ORIGINS"),
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(logger.New(logger.Config{
		// Format: "${time} ${method} ${path} - ${ip} - ${status} - ${latency}\n",
	}))

	// public static files
	// app.Static("/", "./public", fiber.Static{})

	routes.SetupRoutes(app)

	app.Get("/swagger/*", swagger.HandlerDefault)

	fmt.Println("Server is running on http://localhost:" + os.Getenv("PORT"))
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
