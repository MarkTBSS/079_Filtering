package service

import (
	_models "github.com/MarkTBSS/079_Filtering/pkg/itemShop/models"
	"github.com/MarkTBSS/079_Filtering/pkg/itemShop/repository"
)

type itemShopServiceImpl struct {
	itemShopRepository repository.ItemShopRepository
}

func (s *itemShopServiceImpl) Listing(itemFilter *_models.ItemFilter) ([]*_models.Item, error) {
	itemEntityList, err := s.itemShopRepository.Listing(itemFilter)
	if err != nil {
		return nil, err
	}

	itemModelLists := make([]*_models.Item, 0)
	for _, itemEntity := range itemEntityList {
		itemModelLists = append(itemModelLists, itemEntity.ChangeToItemModel())
	}

	return itemModelLists, nil
}

func NewItemShopRepositoryImpl(itemShopRepository repository.ItemShopRepository) ItemShopService {
	return &itemShopServiceImpl{itemShopRepository}
}
