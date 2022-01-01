package main

import (
	"context"
	"ent-go-demo/ent"
	_ "ent-go-demo/ent/runtime"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
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

func mapProject(project *ent.Project) map[string]interface{} {
	return map[string]interface{}{
		"id":      project.ID,
		"name":    project.Name,
		"user_id": project.UserID,
	}
}

func mapProjects(projects []*ent.Project) []map[string]interface{} {
	pp := make([]map[string]interface{}, len(projects))

	for i, p := range projects {
		pp[i] = mapProject(p)
	}

	return pp
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
	app.Use(logger.New())

	v1 := app.Group("/api/v1")

	auth := v1.Group("/auth")

	auth.Post("/register", server.Register)
	auth.Post("/login", server.Login)
	auth.Post("/logout", Protected(), server.Logout)

	v1.Get("/projects", Protected(), server.GetProjects)
	v1.Post("/projects", Protected(), server.CreateProject)

	v1.Get("/me", Protected(), server.Me)

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
