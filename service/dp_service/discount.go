package dp_service

import (
	"github.com/HerryZeng/dpdemo1/model"
	"github.com/HerryZeng/dpdemo1/repository"
)

type DiscountService struct {
	Repo *repository.Discount
}

func (d *DiscountService) ListDiscounts() []model.Discount {
	itemsLeft := d.Repo.ListDiscounts(1)
	itemsRight := d.Repo.ListDiscounts(2)
	itemsLeft = append(itemsLeft, itemsRight...)
	return itemsLeft
}
