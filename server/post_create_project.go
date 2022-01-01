package main

import (
	"ent-go-demo/ent"
	"log"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) CreateProject(c *fiber.Ctx) error {
	type CreateProjectInput struct {
		Name        string  `json:"name"`
		Description *string `json:"description"`
	}

	userID, err := getUserID(c)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	var input CreateProjectInput

	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	q := s.DB.Project.Create().SetUserID(userID).SetName(input.Name)

	if input.Description != nil {
		q.SetDescription(*input.Description)
	}

	p, err := q.Save(c.Context())
	if err != nil {
		log.Println(err)
		if ent.IsConstraintError(err) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "project with this name already exist"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	return c.JSON(fiber.Map{
		"data": p,
	})
}
