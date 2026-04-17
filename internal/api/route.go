package api

import "github.com/gofiber/fiber/v2"

func (s *Server) Route(route fiber.Router) {
	route.Post("/item", s.CreateItem)
	route.Get("/items", s.GetItems)
	route.Delete("/item", s.DeleteItem)
	route.Post("/find", s.FindItems)
	route.Post("/update", s.UpdateItem)
}
