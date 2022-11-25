package res

import "github.com/HerryZeng/dpdemo1/model"

type TeamResp struct {
	model.TeamDetail
	FoodList    []TeamChooseFoodResp
	ChooseItems []TeamChooseItemResp
}
