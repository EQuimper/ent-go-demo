package main

import (
	"ent-go-demo/ent"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func createJWTToken(u *ent.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = u.Username
	claims["email"] = u.Email
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token.SignedString([]byte("mysecret"))
}

func getUserID(c *fiber.Ctx) (int, error) {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userIDStr, ok := claims["id"].(float64)

	if !ok {
		return 0, fmt.Errorf("can't get claims id")
	}

	userID := int(userIDStr)

	return userID, nil
}
