package main

import (
	"ent-go-demo/ent/project"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) GetProjects(c *fiber.Ctx) error {
	userID, err := getUserID(c)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	projects := s.DB.Project.Query().Where(project.UserIDEQ(userID)).AllX(c.Context())

	return c.JSON(fiber.Map{
		"data": projects,
	})
}
