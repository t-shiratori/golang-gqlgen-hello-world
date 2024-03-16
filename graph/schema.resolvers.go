package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"golang-gqlgen-hello-world/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	//panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
	return &model.Todo{
		ID:   "todo-id-0001",
		Text: input.Text,
		Done: false,
		User: &model.User{
			ID:   input.UserID,
			Name: "user-name-0001",
		},
	}, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	//panic(fmt.Errorf("not implemented: Todos - todos"))
	return []*model.Todo{
		{
			ID:   "todo-id-0001",
			Text: "todo1",
			Done: false,
			User: &model.User{
				ID:   "user-id-0001",
				Name: "user-name-0001",
			},
		},
		{
			ID:   "todo-id-0002",
			Text: "todo2",
			Done: false,
			User: &model.User{
				ID:   "user-id-0002",
				Name: "user-name-0002",
			},
		},
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
