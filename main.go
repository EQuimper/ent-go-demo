package main

import (
	"context"
	"ent-go-demo/ent"
	"fmt"
	"log"

	_ "ent-go-demo/ent/runtime"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "entgo_demo"
)

func main() {
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname),
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

	user, err := client.User.
		Create().
		SetUsername("ryan3").
		SetEmail("ryan3@gmail.com").
		SetPassword("password").
		Save(ctx)
	if err != nil {
		panic(err)
	}

	_, err = client.User.UpdateOne(user).
		SetUsername("ryan5").
		SetPassword("password123").
		Save(ctx)
	if err != nil {
		panic(err)
	}

	_, err = client.Project.
		Create().
		SetName("my project").
		SetDescription("a description").
		SetUser(user).
		Save(ctx)
	if err != nil {
		panic(err)
	}

	projects := user.QueryProjects().AllX(ctx)
	for _, project := range projects {
		log.Println(project.String())
	}
}
