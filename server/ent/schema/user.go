package schema

import (
	"context"
	ent2 "ent-go-demo/ent"
	"ent-go-demo/ent/hook"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"golang.org/x/crypto/bcrypt"
	"log"
	"regexp"
)

var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DefaultMixin{},
		TimeMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique().NotEmpty().MinLen(3).MaxLen(50),
		field.String("email").Match(emailRegexp).Unique(),
		field.Text("password").MinLen(6).NotEmpty().Sensitive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("projects", Project.Type),
	}
}

func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.UserFunc(func(ctx context.Context, mutation *ent2.UserMutation) (ent2.Value, error) {
				log.Println("yo")
				if password, ok := mutation.Password(); ok {
					log.Printf("time to hash password %s \n", password)
					hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
					if err != nil {
						return nil, err
					}

					mutation.SetPassword(string(hash))
				}
				return next.Mutate(ctx, mutation)
			})
		}, ent.OpCreate|ent.OpUpdateOne),
	}
}
