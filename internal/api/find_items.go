package api

import (
	"github.com/gofiber/fiber/v2"
)

type FindRequest struct {
	PartName string `json:"part_name"`
}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type FindResponse struct {
	Items []Item `json:"items"`
}

func (s *Server) FindItems(c *fiber.Ctx) error {
	var req FindRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
	}

	if req.PartName == "" {
		return fiber.NewError(fiber.StatusBadRequest, "part name is required")
	}

	items := s.Repository.Find(c.Context(), req.PartName)

	res := make([]Item, 0, len(items))
	for _, item := range items {
		res = append(
			res,
			Item{
				ID:   item.ID,
				Name: item.Name,
			},
		)
	}

	return c.Status(fiber.StatusOK).JSON(FindResponse{Items: res})
}
