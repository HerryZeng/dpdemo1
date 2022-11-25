package repository

import "github.com/HerryZeng/dpdemo1/model"

type GuessRepo struct {
	DB model.DataBase
}

func (g *GuessRepo) ListGuesses() []model.Guess {
	var guessList []model.Guess
	g.DB.MyDB.Find(&guessList)
	return guessList
}
