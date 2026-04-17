package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type UpdateItemRequest struct {
	ID      string `json:"id"`
	NewName string `json:"part_name"`
}

func (s *Server) UpdateItem(c *fiber.Ctx) error {
	var req UpdateItemRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
	}

	if req.NewName == "" {
		return fiber.NewError(fiber.StatusBadRequest, "new name is required")
	}

	err := s.Repository.Update(c.Context(), req.ID, req.NewName)

	if err != nil {
		log.Errorw("s.Repository.Update", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
