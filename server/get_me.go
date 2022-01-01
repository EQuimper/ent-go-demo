package main

import "github.com/gofiber/fiber/v2"

func (s *Server) Me(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	user, err := s.DB.User.Get(c.Context(), userID)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"data": user,
	})
}
