package dp_service

import (
	"github.com/HerryZeng/dpdemo1/model"
	"github.com/HerryZeng/dpdemo1/repository"
	"github.com/HerryZeng/dpdemo1/res"
)

type ListNavItemService struct {
	Repo *repository.ListNavItemRepo
}

func model2res(nav model.Nav) res.Nav {
	navRes := res.Nav{
		NavId: nav.NavId,
		Src:   nav.Src,
		Title: nav.Title,
	}
	return navRes
}

func (i *ListNavItemService) ListNavItems(level int) []res.Nav {
	result := i.Repo.ListNavItems(level)
	var newList []res.Nav
	for _, item := range result {
		r := model2res(item)
		newList = append(newList, r)
	}

	return newList
}
