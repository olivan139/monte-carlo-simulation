package simulation

import (
	"fmt"
	"monte-carlo-simulation/matrix"
	"monte-carlo-simulation/model"
	"sync"
)

type Result struct {
	TotalWin        int
	TotalBet        int
	MultSymbolCount int
	MultValueCount  int
	FreeSpinsCount  int
}

func Start() {
	ch := make(chan Result, model.Model.NumOfIterations)
	wg := sync.WaitGroup{}
	spinBet := 1
	var finalResult Result

	for i := 0; i < model.Model.NumOfIterations; i++ {
		go countSpin(ch, spinBet, &wg)
	}

	wg.Wait()

	close(ch)

	for result := range ch {
		finalResult.FreeSpinsCount += result.FreeSpinsCount
		finalResult.MultSymbolCount += result.MultSymbolCount
		finalResult.MultValueCount += result.MultValueCount
		finalResult.TotalBet += result.TotalBet
		finalResult.TotalWin += result.TotalWin
	}

	RTPVal := float32(finalResult.TotalWin) / float32(finalResult.TotalBet)
	avgMultVal := float32(finalResult.MultValueCount) / float32(finalResult.FreeSpinsCount+1)
	avgFreeGamesVal := float32(finalResult.FreeSpinsCount) / float32(model.Model.NumOfIterations)

	fmt.Printf("Number of iterations: %v mln.\nRTP: %v\nAVG mult value per free games: %v\nAVG number of free games: %v\n", float32(model.Model.NumOfIterations)/1000000., RTPVal, avgMultVal, avgFreeGamesVal)
}

func countSpin(ch chan Result, spinBet int, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	var result Result
	var mat matrix.Matrix

	mat.Init(3, 5)
	mat.GenerateFromReels(model.Model.Reels)

	for _, line := range model.Model.Lines {
		winLine := []int{mat.Matrix[line[0]][0], mat.Matrix[line[1]][1], mat.Matrix[line[2]][2], mat.Matrix[line[3]][3], mat.Matrix[line[4]][4]}
		result.TotalWin += matrix.GetLinePayoff(winLine)
	}

	result.TotalWin += mat.GetScatterPayoff()
	numOfFreeGames := mat.CheckForFreeGames()

	spinWinAdd := 0
	spinWinAdd, result.MultSymbolCount, result.MultValueCount, result.FreeSpinsCount = countSpinWinForFreeGames(numOfFreeGames)
	result.TotalWin += spinWinAdd
	result.TotalBet += spinBet

	ch <- result
}

func countSpinWinForFreeGames(numOfFreeGames int) (int, int, int, int) {
	var mat matrix.Matrix
	mat.Init(3, 5)

	spinWin := 0
	multiply := 2
	multSymbolCount := 0
	multValueCount := 0
	mreeSpinsCount := 0

	for numOfFreeGames > 0 {
		numOfFreeGames--
		mat.GenerateFromReels(model.Model.FreeGamesReels)
		multiply += mat.GetMultiplierCount()
		if mat.GetMultiplierCount() != 0 {
			multSymbolCount += 1
		}

		multValueCount += multiply
		mreeSpinsCount += 1

		for _, line := range model.Model.Lines {
			winLine := []int{mat.Matrix[line[0]][0], mat.Matrix[line[1]][1], mat.Matrix[line[2]][2], mat.Matrix[line[3]][3], mat.Matrix[line[4]][4]}
			spinWin += multiply * matrix.GetLinePayoff(winLine)
		}

		spinWin += multiply * mat.GetScatterPayoff()
	}

	return spinWin, multSymbolCount, multValueCount, mreeSpinsCount
}
