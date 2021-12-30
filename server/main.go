package main

import (
	"context"
	"ent-go-demo/ent"
	"ent-go-demo/ent/project"
	_ "ent-go-demo/ent/runtime"
	"ent-go-demo/ent/user"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

const (
	host     = "localhost"
	port     = 5432
	dbUser   = "postgres"
	password = "postgres"
	dbname   = "entgo_demo"
)

type Server struct {
	DB *ent.Client
}

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

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "logged in",
		"data": map[string]interface{}{
			"user":         u,
			"access_token": t,
		},
	})
}

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

	cookie := new(fiber.Cookie)
	cookie.Name = "authorization"
	cookie.Value = t
	cookie.Expires = time.Now().Add(time.Hour * 72)
	cookie.HTTPOnly = true

	c.Cookie(cookie)

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "user created",
		"data": map[string]interface{}{
			"user":         u,
			"access_token": t,
		},
	})
}

func createJWTToken(u *ent.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = u.Username
	claims["email"] = u.Email
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token.SignedString([]byte("mysecret"))
}

func getUserID(c *fiber.Ctx) (uuid.UUID, error) {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userIDStr := claims["id"].(string)

	return uuid.Parse(userIDStr)
}

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

func main() {
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, dbUser, password, dbname),
	)
	if err != nil {
		log.Fatalf("failed connecting to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run migration.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	server := Server{DB: client}

	app := fiber.New()

	app.Use(cors.New())

	v1 := app.Group("/api/v1")

	auth := v1.Group("/auth")

	auth.Post("/register", server.Register)
	auth.Post("/login", server.Login)

	v1.Get("/projects", Protected(), server.GetProjects)
	v1.Post("/projects", Protected(), server.CreateProject)

	log.Fatal(app.Listen(":4000"))
	//
	//user, err := client.User.
	//	Create().
	//	SetUsername("ryan3").
	//	SetEmail("ryan3@gmail.com").
	//	SetPassword("password").
	//	Save(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//
	//_, err = client.User.UpdateOne(user).
	//	SetUsername("ryan5").
	//	SetPassword("password123").
	//	Save(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//
	//_, err = client.Project.
	//	Create().
	//	SetName("my project").
	//	SetDescription("a description").
	//	SetUser(user).
	//	Save(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//
	//_, err = client.Project.
	//	Create().
	//	SetName("my second project").
	//	SetDescription("a description").
	//	SetUser(user).
	//	Save(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//
	//projects := user.QueryProjects().Order(ent.Asc(project.FieldCreatedAt)).AllX(ctx)
	//for _, p := range projects {
	//	log.Println(p.String())
	//}
}
