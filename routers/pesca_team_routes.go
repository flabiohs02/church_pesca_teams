package routers

import (
	"church_pesca_teams/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupPescaTeamRoutes(app *fiber.App, ptHandler *handler.PescaTeamHandler) {
	v1 := app.Group("/api/v1")
	{
		v1.Post("/pesca-teams", ptHandler.CreatePescaTeam)
		v1.Get("/pesca-teams/:id", ptHandler.GetPescaTeamByID)
		v1.Get("/pesca-teams", ptHandler.GetAllPescaTeams)
		v1.Put("/pesca-teams/:id", ptHandler.UpdatePescaTeam)
		v1.Delete("/pesca-teams/:id", ptHandler.DeletePescaTeam)
	}
}
