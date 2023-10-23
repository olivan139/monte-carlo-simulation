package matrix

import (
	"hash/maphash"

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

func (m *Matrix) GenerateFromReels(reels [][]int) {
	var randIndArr []int

	for i := range reels {
		outUint64 := new(maphash.Hash).Sum64()
		out := int(outUint64)
		if out < 0 {
			out = -out
		}
		index := out % len(reels[i])
		randIndArr = append(randIndArr, index)
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
