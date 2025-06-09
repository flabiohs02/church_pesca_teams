package handler

import (
	"strconv"

	"church_pesca_teams/domain"
	"church_pesca_teams/usecase"

	"github.com/gofiber/fiber/v2"
)

type PescaTeamHandler struct {
	ptService *usecase.PescaTeamService
}

func NewPescaTeamHandler(service *usecase.PescaTeamService) *PescaTeamHandler {
	return &PescaTeamHandler{ptService: service}
}

func (h *PescaTeamHandler) CreatePescaTeam(c *fiber.Ctx) error {
	var pt domain.PescaTeam
	if err := c.BodyParser(&pt); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.ptService.CreatePescaTeam(&pt); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(pt)
}

func (h *PescaTeamHandler) GetPescaTeamByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	pt, err := h.ptService.GetPescaTeamByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Pesca team not found"})
	}

	return c.Status(fiber.StatusOK).JSON(pt)
}

func (h *PescaTeamHandler) GetAllPescaTeams(c *fiber.Ctx) error {
	pts, err := h.ptService.GetAllPescaTeams()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(pts)
}

func (h *PescaTeamHandler) UpdatePescaTeam(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var pt domain.PescaTeam
	if err := c.BodyParser(&pt); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	pt.ID = uint(id)
	if err := h.ptService.UpdatePescaTeam(&pt); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(pt)
}

func (h *PescaTeamHandler) DeletePescaTeam(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.ptService.DeletePescaTeam(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
