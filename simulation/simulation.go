package simulation

import (
	"fmt"
	"monte-carlo-simulation/matrix"
	"monte-carlo-simulation/model"
)

func StartSimulation(iter int) {
	m := 0
	c := 0
	o := 0
	totalBet := 0
	totalWin := 0

	spinBet := 1
	var mat matrix.Matrix
	mat.Init(3, 5)

	for i := 0; i < iter; i++ {
		totalBet += spinBet
		mat.GenerateFromReels(model.Model.Reels)
		spinWin := 0
		for _, line := range model.Model.Lines {
			winLine := []int{mat.Matrix[line[0]][0], mat.Matrix[line[1]][1], mat.Matrix[line[2]][2], mat.Matrix[line[3]][3], mat.Matrix[line[4]][4]}
			spinWin += matrix.GetLinePayoff(winLine)
		}

		spinWin += mat.GetScatterPayoff()
		numOfFreeGames := mat.CheckForFreeGames()
		spinWinAdd, o1, m1, c1 := countSpinWinForFreeGames(numOfFreeGames)

		totalWin += spinWin + spinWinAdd
		o += o1
		m += m1
		c += c1
		if i%100000 == 0 {
			fmt.Println(float32(totalWin)/float32(totalBet)*100.0, i, float32(m)/float32(c+1), float32(c)/float32(i+1), float32(m)/float32(i+1), float32(o)/float32(c+1), float32(o)/float32(i+1))
		}
	}

}

func countSpinWinForFreeGames(numOfFreeGames int) (int, int, int, int) {
	var mat matrix.Matrix
	mat.Init(3, 5)
	spinWin := 0
	multiply := 2
	o := 0
	m := 0
	c := 0
	for numOfFreeGames > 0 {
		numOfFreeGames--
		mat.GenerateFromReels(model.Model.FreeGamesReels)
		multiply += mat.GetMultiplierCount()
		if mat.GetMultiplierCount() != 0 {
			o += 1
		}

		m += multiply
		c += 1

		for _, line := range model.Model.Lines {
			winLine := []int{mat.Matrix[line[0]][0], mat.Matrix[line[1]][1], mat.Matrix[line[2]][2], mat.Matrix[line[3]][3], mat.Matrix[line[4]][4]}
			spinWin += multiply * matrix.GetLinePayoff(winLine)
		}

		spinWin += multiply * mat.GetScatterPayoff()
	}

	return spinWin, o, m, c
}
