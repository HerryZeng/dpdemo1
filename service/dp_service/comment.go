package dp_service

import (
	"github.com/HerryZeng/dpdemo1/model"
	"github.com/HerryZeng/dpdemo1/repository"
)

type CommentService struct {
	Repo repository.CommentRepo
}

func (c *CommentService) GetCommentList(hotelId string) []model.Comment {
	return c.Repo.GetCommentList(hotelId)
}
