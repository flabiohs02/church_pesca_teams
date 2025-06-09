package main

import (
	"log"

	"church_pesca_teams/config"
	"church_pesca_teams/domain"
	"church_pesca_teams/handler"
	"church_pesca_teams/repository"
	"church_pesca_teams/routers"
	"church_pesca_teams/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	db := config.InitDb()

	// Auto-migrate the PescaTeam schema
	db.AutoMigrate(&domain.PescaTeam{})

	// Dependency Injection
	ptRepo := repository.NewGormPescaTeamRepository(db)
	ptService := usecase.NewPescaTeamService(ptRepo)
	ptHandler := handler.NewPescaTeamHandler(ptService)

	// Set up Fiber app
	app := fiber.New()

	// Setup routes
	routers.SetupPescaTeamRoutes(app, ptHandler)

	log.Println("Server listening on :8787")
	if err := app.Listen(":8787"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
