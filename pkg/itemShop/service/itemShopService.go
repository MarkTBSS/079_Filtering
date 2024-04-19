package service

import _models "github.com/MarkTBSS/079_Filtering/pkg/itemShop/models"

type ItemShopService interface {
	Listing(itemFilter *_models.ItemFilter) ([]*_models.Item, error)
}
