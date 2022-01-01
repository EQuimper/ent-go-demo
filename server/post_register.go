package main

import (
	"ent-go-demo/ent/user"
	"log"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) Register(c *fiber.Ctx) error {
	ctx := c.Context()

	type RegisterInput struct {
		Username             string `json:"username"`
		Email                string `json:"email"`
		Password             string `json:"password"`
		PasswordConfirmation string `json:"password_confirmation"`
	}

	var registerForm RegisterInput

	err := c.BodyParser(&registerForm)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on register request", "data": err})
	}

	if registerForm.PasswordConfirmation != registerForm.Password {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "password_confirmation must match password", "data": err})
	}

	exist := s.DB.User.Query().Where(user.Or(
		user.UsernameEQ(registerForm.Username),
		user.EmailEQ(registerForm.Email),
	)).ExistX(ctx)
	if exist {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "User with this username or email already exist", "data": err})
	}

	u, err := s.DB.User.
		Create().
		SetUsername(registerForm.Username).
		SetEmail(registerForm.Email).
		SetPassword(registerForm.Password).
		Save(ctx)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error register", "data": err})
	}

	t, err := createJWTToken(u)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	setAuthCookie(c, t)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "user created",
		"data": map[string]interface{}{
			"user":         u,
			"access_token": t,
		},
	})
}
