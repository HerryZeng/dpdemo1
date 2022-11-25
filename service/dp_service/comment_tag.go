package dp_service

import (
	"github.com/HerryZeng/dpdemo1/model"
	"github.com/HerryZeng/dpdemo1/repository"
)

type CommentTagService struct {
	Repo repository.CommentTagRepo
}

func (c *CommentTagService) GetCommentTagList(hotelId string) []model.CommentTag {
	return c.Repo.GetCommentTagList(hotelId)
}
