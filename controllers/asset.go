package controllers

import (
	"context"
	"github.com/jterrazz/ccap.live-api/models"
	"github.com/jterrazz/ccap.live-api/services"
)

func BTAssetsToAssets(original []*services.BTAsset) []*models.Asset {
	var assets []*models.Asset

	for _, b := range original {
		newFoo := models.Asset{Name: b.AssetName}
		assets = append(assets, &newFoo)
	}

	return assets
}

func (r *queryResolver) Assets (ctx context.Context) ([]*models.Asset, error) {
	assets := services.GetAssets()
	return BTAssetsToAssets(assets), nil
}