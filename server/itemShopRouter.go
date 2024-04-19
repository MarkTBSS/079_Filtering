package server

import (
	"github.com/MarkTBSS/079_Filtering/pkg/itemShop/controller"
	"github.com/MarkTBSS/079_Filtering/pkg/itemShop/repository"
	"github.com/MarkTBSS/079_Filtering/pkg/itemShop/service"
)

func (s *echoServer) initItemShopRouter() {
	router := s.app.Group("/v1/item-shop")

	itemShopRepository := repository.NewItemShopRepositoryImpl(s.db, s.app.Logger)
	itemShopService := service.NewItemShopRepositoryImpl(itemShopRepository)
	itemShopController := controller.NewItemShopControllerImpl(itemShopService)

	router.GET("/listing", itemShopController.Listing)
}
