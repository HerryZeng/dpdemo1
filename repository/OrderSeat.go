package repository

import (
	"log"

	"github.com/HerryZeng/dpdemo1/model"
)

type OrderSeatRepo struct {
	DB model.DataBase
}

func (s *OrderSeatRepo) OrderSeatOp(o model.OrderSeat) string {
	err := s.DB.MyDB.Save(o).Error
	if err != nil {
		log.Println(err)
	}
	return o.OrderSeatId
}
