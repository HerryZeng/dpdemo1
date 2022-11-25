package dp_service

import (
	"github.com/HerryZeng/dpdemo1/model"
	"github.com/HerryZeng/dpdemo1/repository"
)

type MarketService struct {
	Repo repository.MarketRepo
}

func (m *MarketService) GetMarketInfo(hotelId string) model.Market {
	var market model.Market
	m.Repo.DB.MyDB.Where("hotel_id=?", hotelId).Find(&market)
	return market
}
