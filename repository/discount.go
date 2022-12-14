package repository

import (
	"github.com/HerryZeng/dpdemo1/model"
)

type Discount struct {
	DB model.DataBase
}

func (d *Discount) ListDiscounts(pos int) []model.Discount {
	var discounts []model.Discount
	d.DB.MyDB.Where("pos=?", pos).Find(&discounts)
	return discounts
}
