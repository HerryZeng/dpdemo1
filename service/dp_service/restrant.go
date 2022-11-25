package dp_service

import (
	"github.com/HerryZeng/dpdemo1/model"
	"github.com/HerryZeng/dpdemo1/repository"
)

type RestaurantService struct {
	Repo     *repository.RestaurantNavRepo
	ItemRepo *repository.RestaurantTabItemRepo
}

func (r *RestaurantService) ListRestaurantNav(level int) []model.RestaurantNav {
	items := r.Repo.ListRestaurantNav(level)
	return items
}

func (r *RestaurantService) ListGoodRestaurantTabItem() []model.RestaurantTabItem {
	items := r.ItemRepo.ListGoodRestaurantTabItem()
	return items
}
