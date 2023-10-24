package simulation

import (
	"fmt"
	"monte-carlo-simulation/matrix"
	"monte-carlo-simulation/model"
	"sync"
	"sync/atomic"
)

var WinSum atomic.Int64

func Start() {
	wg := sync.WaitGroup{}
	spinBet := 1

	for i := 0; i < model.Model.NumOfIterations; i++ {
		go countSpin(spinBet, &wg)
	}

	wg.Wait()

	RTPVal := float32(WinSum.Load()) / float32(model.Model.NumOfIterations)

	fmt.Printf("Number of iterations: %v mln.\nRTP: %v\n", float32(model.Model.NumOfIterations)/1000000., RTPVal)
}

func countSpin(spinBet int, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	var mat matrix.Matrix
	result := 0
	mat.Init(3, 5)
	mat.GenerateFromReels(model.Model.Reels)

	for _, line := range model.Model.Lines {
		winLine := []int{mat.Matrix[line[0]][0], mat.Matrix[line[1]][1], mat.Matrix[line[2]][2], mat.Matrix[line[3]][3], mat.Matrix[line[4]][4]}
		result += matrix.GetLinePayoff(winLine)
	}

	result += mat.GetScatterPayoff()
	numOfFreeGames := mat.CheckForFreeGames()
	WinSum.Add(int64(result))

	wg.Add(1)
	go countSpinWinForFreeGames(numOfFreeGames, wg)
}

func countSpinWinForFreeGames(numOfFreeGames int, wg *sync.WaitGroup) {
	defer wg.Done()
	var multiplier atomic.Int64
	multiplier.Add(2)
	for ; numOfFreeGames > 0; numOfFreeGames-- {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var mat matrix.Matrix
			mat.Init(3, 5)
			spinWin := 0
			mat.GenerateFromReels(model.Model.FreeGamesReels)
			multiplier.Add(int64(mat.GetMultiplierCount()))

			for _, line := range model.Model.Lines {
				winLine := []int{mat.Matrix[line[0]][0], mat.Matrix[line[1]][1], mat.Matrix[line[2]][2], mat.Matrix[line[3]][3], mat.Matrix[line[4]][4]}
				spinWin += int(multiplier.Load()) * matrix.GetLinePayoff(winLine)
			}

			spinWin += int(multiplier.Load()) * mat.GetScatterPayoff()
			WinSum.Add(int64(spinWin))

		}()
	}
}
