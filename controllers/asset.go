package controllers

import (
	"context"
	"github.com/jterrazz/ccap.live-api/models"
	"github.com/jterrazz/ccap.live-api/services"
)

func BTAssetsToAssets(original []*services.BTAsset) []*models.Asset {
	var assets []*models.Asset

	for _, b := range original {
		newAsset := models.Asset{Name: b.AssetName, }
		assets = append(assets, &newAsset)
	}

	return assets
}

func (r *queryResolver) Assets (ctx context.Context) ([]*models.Asset, error) {
	assets, err := services.GetAssets()
	if err != nil {
		return nil, err
	}
	return BTAssetsToAssets(assets), nil
}