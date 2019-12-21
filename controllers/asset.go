package controllers

import (
	"context"
	"github.com/jterrazz/ccap.live-api/models"
	"github.com/jterrazz/ccap.live-api/services"
)

func (r *queryResolver) Assets (ctx context.Context) ([]*models.Asset, error) {
	assets := services.GetAssets()
	return assets, nil
}