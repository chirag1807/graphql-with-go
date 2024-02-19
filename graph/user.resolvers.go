package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"articlewithgraphql/api/repository"
	errorhandling "articlewithgraphql/error"
	"articlewithgraphql/graph/model"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// RegisterUser is the resolver for the registerUser field.
func (r *mutationResolver) RegisterUser(ctx context.Context, input model.RegisterUser) (string, error) {
	var user model.RegisterUser
	user.Name = input.Name
	user.Bio = input.Bio
	user.Email = input.Email
	user.Password = input.Password
	message, err := repository.RegisterUser(ctx.Value("conn").(*pgxpool.Pool), user)

	if err != nil {
		return "", err
		// return utils.ErrorGenerator(err)
	} else {
		return message, err
	}
}

// LoginUser is the resolver for the loginUser field.
func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginUser) (*model.User, error) {
	fmt.Println(ctx.Value("id"))
	var user model.LoginUser
	user.Email = input.Email
	user.Password = input.Password
	dbUser, err := repository.LoginUser(ctx.Value("conn").(*pgxpool.Pool), user)
	return &dbUser, err
}

// AddArticle is the resolver for the addArticle field.
func (r *mutationResolver) AddArticle(ctx context.Context, input model.AddArticle) (string, error) {
	var article = input
	userID, _ := ctx.Value("id").(int64)
	fmt.Println(userID)
	article.Author = &userID
	message, err := repository.AddArticle(ctx.Value("conn").(*pgxpool.Pool), article)
	return message, err
}

// UpdateArticle is the resolver for the updateArticle field.
func (r *mutationResolver) UpdateArticle(ctx context.Context, input model.UpdateArticle) (string, error) {
	if ctx.Value("id") == nil {
		return "", errorhandling.TokenNotFound
	}
	var article = input
	message, err := repository.UpdateArticle(ctx.Value("conn").(*pgxpool.Pool), article)
	return message, err
}

// AddTopic is the resolver for the addTopic field.
func (r *mutationResolver) AddTopic(ctx context.Context, input model.AddTopic) (string, error) {
	var topic = input
	message, err := repository.AddTopic(ctx.Value("conn").(*pgxpool.Pool), topic)
	return message, err
}

// UpdateTopic is the resolver for the updateTopic field.
func (r *mutationResolver) UpdateTopic(ctx context.Context, input model.UpdateTopic) (string, error) {
	var topic = input
	message, err := repository.UpdateTopic(ctx.Value("conn").(*pgxpool.Pool), topic)
	return message, err
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
