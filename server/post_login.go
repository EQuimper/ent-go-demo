package main

import (
	"ent-go-demo/ent/user"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var input LoginInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}

	u, err := s.DB.User.Query().Where(user.EmailEQ(input.Email)).First(c.Context())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "bad email/password combination"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(input.Password)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "bad email/password combination"})
	}

	t, err := createJWTToken(u)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	setAuthCookie(c, t)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "logged in",
		"data": map[string]interface{}{
			"user":         u,
			"access_token": t,
		},
	})
}
