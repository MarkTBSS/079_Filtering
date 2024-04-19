package repository

import (
	_entities "github.com/MarkTBSS/079_Filtering/entities"
	_models "github.com/MarkTBSS/079_Filtering/pkg/itemShop/models"
)

type ItemShopRepository interface {
	Listing(itemFilter *_models.ItemFilter) ([]*_entities.Item, error)
}
