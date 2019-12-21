package server

import (
	"context"
	models "github.com/jterrazz/ccap.live-api/models"
)

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, email string) (bool, error) {
	return true, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	users := []*models.User{&models.User{Email: "emmm", ID: "42", Payload: "OK"}}
	return users, nil
}
