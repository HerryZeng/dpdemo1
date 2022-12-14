package repository

import "github.com/HerryZeng/dpdemo1/model"

type HotelRepo struct {
	DB model.DataBase
}

func (h *HotelRepo) GetHotelById(hotelId string) model.Hotel {
	var hotel model.Hotel
	//select * from hotel where hotel_id = hotelId
	h.DB.MyDB.Where("hotel_id = ?", hotelId).First(&hotel)
	return hotel
}
