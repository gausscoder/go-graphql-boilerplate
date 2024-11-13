package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"context"
	"go-graphql-boilerplate/datasources/mutation"
	"go-graphql-boilerplate/datasources/query"
	"go-graphql-boilerplate/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, firstName string, lastName string) (*model.CreateUserResponse, error) {
	response, _ := mutation.CreateUserHandler(ctx, firstName, lastName)
	return response, nil
}

// GetAllUsers is the resolver for the getAllUsers field.
func (r *queryResolver) GetAllUsers(ctx context.Context) (*model.GetAllUsersResponse, error) {
	response, _ := query.GetAllUsersHandler(ctx)
	return response, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
