package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func setAuthCookie(c *fiber.Ctx, token string) {
	cookie := new(fiber.Cookie)
	cookie.Name = "authorization"
	cookie.Value = token
	cookie.Expires = time.Now().Add(time.Hour * 72)
	cookie.HTTPOnly = true
	cookie.SameSite = "lax"
	cookie.Domain = "localhost"

	c.Cookie(cookie)
}
