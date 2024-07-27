package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
	"go-gqlgen/graph/model"
)

// CreateUser is the resolver for the CreateUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input *model.UserInput) (*model.ResponseCreateUser, error) {
	return &model.ResponseCreateUser{
		Name:  input.Name,
		Setor: input.Setor,
		Message: "user created successfuly!",
	}, nil
}

// GetUser is the resolver for the GetUser field.
func (r *queryResolver) GetUser(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented: GetUser - GetUser"))
}
