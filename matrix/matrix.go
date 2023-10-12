package matrix

import (
	"crypto/rand"
	"log"
	"math/big"

	"monte-carlo-simulation/helper"
	"monte-carlo-simulation/model"
)

type Matrix struct {
	Matrix [][]int
	Rows   int
	Cols   int
}

func (m *Matrix) Init(rows int, cols int) {
	m.Matrix = make([][]int, rows)
	m.Cols = cols
	m.Rows = rows

	for i := range m.Matrix {
		m.Matrix[i] = make([]int, cols)
	}
}

func (m *Matrix) GenerateFromReels(reels [][]int) error {
	var randIndArr []int

	for i := range reels {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(reels[i]))))
		if err != nil {
			log.Panic(err)
			return err
		}

		randIndArr = append(randIndArr, int(index.Uint64()))
	}

	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			m.Matrix[i][j] = reels[j][randIndArr[j]]
			randIndArr[j]++

			if randIndArr[j] == len(reels[j]) {
				randIndArr[j] = 0
			}
		}
	}

	return nil
}

func (m *Matrix) CheckForFreeGames() int {
	transposedMatrix := helper.Transpose(m.Matrix)
	numOfScatters := 0
	for i := range transposedMatrix {
		for j := range transposedMatrix[i] {
			if model.Model.ScatterSymbol == transposedMatrix[i][j] {
				numOfScatters++
				break
			}
		}
	}

	return model.Model.FreeGames[numOfScatters]
}

func (m *Matrix) GetScatterPayoff() int {
	transposedMatrix := helper.Transpose(m.Matrix)
	numOfScatters := 0
	for i := range transposedMatrix {
		for j := range transposedMatrix[i] {
			if model.Model.ScatterSymbol == transposedMatrix[i][j] {
				numOfScatters++
				break
			}
		}
	}

	return model.Model.Paytable[model.Model.ScatterSymbol][numOfScatters]
}

func GetLinePayoff(winLine []int) int {

	frstSymbol := winLine[0]
	wildCount := 0
	wildCountAsLine := 0
	wildAsLine := false
	symbolCount := 0
	mainSymbol := -1

	if frstSymbol == model.Model.WildSymbol {
		wildCount++
		wildCountAsLine++
		wildAsLine = true
	} else {
		symbolCount++
		mainSymbol = frstSymbol
	}

	for i := 1; i < len(winLine); i++ {
		if winLine[i] != model.Model.WildSymbol {
			wildAsLine = false

			if mainSymbol == -1 {
				mainSymbol = winLine[i]
				symbolCount++
			} else if mainSymbol == winLine[i] {
				symbolCount++
			} else {
				break
			}
		} else {
			if wildAsLine {
				wildCountAsLine++
			}

			wildCount++
		}
	}

	if mainSymbol == -1 {
		return model.Model.Paytable[model.Model.WildSymbol][wildCountAsLine]
	}
	if mainSymbol == model.Model.ScatterSymbol {
		return model.Model.Paytable[model.Model.ScatterSymbol][symbolCount]
	}
	return helper.Max(model.Model.Paytable[mainSymbol][symbolCount+wildCount],
		model.Model.Paytable[model.Model.WildSymbol][wildCountAsLine])
}

func (m *Matrix) GetMultiplierCount() int {
	transposedMatrix := helper.Transpose(m.Matrix)
	counter := 0
	for i := range transposedMatrix[model.Model.MultiplyReelNum] {
		if transposedMatrix[model.Model.MultiplyReelNum][i] == model.Model.MultiplySymbol {
			counter++
		}
	}

	return counter
}
