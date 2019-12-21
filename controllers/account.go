package controllers

import (
	"context"

	"github.com/jterrazz/ccap.live-api/models"
)

func (r *mutationResolver) CreateAccount(ctx context.Context, email string) (bool, error) {
	return true, nil
}

func (r *mutationResolver) UpdateAccount(ctx context.Context, email string, token string, payload string) (bool, error) {
	return true, nil
}

func (r *queryResolver) Account(ctx context.Context, email string) (*models.Account, error) {
	return &models.Account{Email: "emmm", ID: "42", Payload: nil}, nil
}