package dp_service

import (
	"github.com/HerryZeng/dpdemo1/model"
	"github.com/HerryZeng/dpdemo1/repository"
)

type MeService struct {
	Repo *repository.MeItemRepo
}

func (m *MeService) ListMe() []model.MeItem {
	return m.Repo.ListMe()
}
