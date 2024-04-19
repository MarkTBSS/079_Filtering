package repository

import (
	"github.com/MarkTBSS/079_Filtering/databases"
	_entities "github.com/MarkTBSS/079_Filtering/entities"
	_exception "github.com/MarkTBSS/079_Filtering/pkg/itemShop/exeption"
	_models "github.com/MarkTBSS/079_Filtering/pkg/itemShop/models"
	"github.com/labstack/echo/v4"
)

type itemRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func (r *itemRepositoryImpl) Listing(itemFilter *_models.ItemFilter) ([]*_entities.Item, error) {
	query := r.db.Connect().Model(&_entities.Item{})
	itemLists := make([]*_entities.Item, 0)

	if itemFilter.Name != "" {
		query = query.Where("name ilike ?", "%"+itemFilter.Name+"%")
	}
	if itemFilter.Description != "" {
		query = query.Where("description ilike ?", "%"+itemFilter.Description+"%")
	}

	err := query.Find(&itemLists).Error
	if err != nil {
		r.logger.Error("Failed to find items", err.Error())
		return nil, &_exception.ItemListing{}
	}
	return itemLists, nil
	// Error Testing
	//return nil, &_exception.ItemListing{}

}
func NewItemShopRepositoryImpl(db databases.Database, logger echo.Logger) ItemShopRepository {
	return &itemRepositoryImpl{
		db:     db,
		logger: logger,
	}
}
