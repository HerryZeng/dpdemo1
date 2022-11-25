package dp_service

import (
	"github.com/HerryZeng/dpdemo1/model"
	"github.com/HerryZeng/dpdemo1/repository"
)

type GuessService struct {
	Repo *repository.GuessRepo
}

func (g *GuessService) ListGuess() []model.Guess {
	return g.Repo.ListGuesses()
}
