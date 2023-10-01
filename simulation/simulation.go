package simulation

import (
	"fmt"
	"monte-carlo-simulation/gameDescription"
	"monte-carlo-simulation/matrix"
)

func StartSimulation() {
	m := 0
	c := 0
	o := 0
	cycle := 10000000000
	totalBet := 0
	totalWin := 0
	reels := [][]int{
		{6, 8, 8, 3, 4, 5, 2, 11, 5, 4, 2, 3, 2, 4, 6, 5, 2, 6, 8, 6, 5, 7, 4, 9, 2, 11, 3, 3, 6, 2, 6, 5, 3, 4, 7, 9, 4, 10, 5, 8, 2, 3, 7, 3, 10, 11, 2, 2, 4, 5, 7, 2, 10, 2, 3, 4, 4, 7, 3, 9, 4, 5, 6, 4, 3, 6, 2, 7, 2, 5, 2, 8, 3, 5, 2, 4, 3, 9, 6, 2, 8, 4, 2, 6, 3, 9, 3},
		{8, 2, 4, 8, 4, 5, 9, 3, 7, 3, 4, 11, 3, 3, 4, 2, 6, 2, 3, 7, 8, 3, 2, 5, 2, 4, 5, 2, 3, 6, 6, 5, 4, 4, 8, 2, 9, 3, 3, 7, 2, 7, 5, 5, 4, 2, 5, 5, 6, 9, 7, 5, 9, 11, 3, 4, 6, 2, 3, 2, 5, 10, 4, 2, 10, 9, 4, 8, 2, 5, 2, 10, 2, 3, 6, 6, 4, 11, 2, 8, 3, 2, 6},
		{10, 1, 3, 7, 5, 11, 9, 1, 2, 6, 4, 8, 10, 7, 4, 4, 8, 6, 3, 9, 5, 8, 3, 2, 4, 2, 3, 3, 5, 2, 4, 5, 5, 1, 3, 11, 6, 2, 1, 2, 4, 6, 2, 9, 2, 8, 2, 5, 4, 3, 6, 4, 4, 2, 5, 7, 4, 9, 3, 4, 4, 7, 6, 6, 10, 3, 8, 7, 3, 3, 5, 3, 5, 3, 3, 7, 3, 6, 5, 10, 2, 4, 2, 2, 4, 11, 2, 2, 5, 6, 4, 2, 9},
		{2, 6, 2, 10, 2, 6, 6, 11, 2, 4, 8, 9, 2, 3, 3, 7, 2, 5, 8, 5, 5, 10, 7, 6, 3, 3, 11, 10, 4, 3, 4, 2, 2, 3, 5, 2, 2, 5, 5, 6, 4, 2, 5, 9, 6, 8, 3, 9, 3, 3, 5, 9, 3, 8, 11, 4, 6, 3, 4, 3, 6, 9, 2, 7, 5, 8, 6, 2, 3, 2, 4, 4, 4, 7, 4, 7, 2, 4, 6, 8, 7, 2, 4, 3, 5, 2, 2},
		{2, 5, 2, 2, 8, 6, 8, 8, 3, 4, 8, 6, 7, 6, 4, 6, 6, 5, 3, 3, 4, 11, 2, 2, 3, 3, 4, 3, 4, 3, 7, 5, 7, 5, 2, 4, 2, 7, 3, 4, 4, 4, 2, 8, 9, 9, 3, 4, 2, 4, 5, 3, 6, 7, 2, 2, 5, 4, 3, 11, 3, 4, 9, 10, 2, 5, 8, 11, 2, 6, 10, 6, 10, 5, 2, 3, 2, 5, 5, 2, 6, 2, 4, 5, 2, 2, 10, 3, 9},
	}
	spinBet := 20
	var mat matrix.Matrix
	mat.Init(3, 5)

	for i := 0; i < cycle; i++ {
		totalBet += spinBet
		mat.GenerateFromReels(reels)
		spinWin := 0
		for _, line := range gameDescription.Desc.Lines {
			winLine := []int{mat.Matrix[line[0]][0], mat.Matrix[line[1]][1], mat.Matrix[line[2]][2], mat.Matrix[line[3]][3], mat.Matrix[line[4]][4]}
			spinWin += matrix.GetLinePayoff(winLine)
		}

		spinWin += mat.GetScatterPayoff()
		numOfFreeGames := mat.CheckForFreeGames()
		multiply := 2
		if numOfFreeGames != 0 {
			for numOfFreeGames > 0 {
				numOfFreeGames--
				mat.GenerateFromReels(reels)
				multiply += mat.GetMultiplierCount()
				if mat.GetMultiplierCount() != 0 {
					o += 1
				}

				m += multiply
				c += 1

				for _, line := range gameDescription.Desc.Lines {
					winLine := []int{mat.Matrix[line[0]][0], mat.Matrix[line[1]][1], mat.Matrix[line[2]][2], mat.Matrix[line[3]][3], mat.Matrix[line[4]][4]}
					spinWin += multiply * matrix.GetLinePayoff(winLine)
				}

				spinWin += multiply * mat.GetScatterPayoff()
			}
		}

		totalWin += spinWin
		if i%100000 == 0 {
			fmt.Println(float32(totalWin)/float32(totalBet)*100.0, i, m/(c+1), c/(i+1), m/(i+1), o/(c+1), o/(i+1))
		}
	}

}
