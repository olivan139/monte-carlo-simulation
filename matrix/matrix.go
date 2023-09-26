package matrix

import (
	"crypto/rand"
	"log"
	"math/big"
	"monte-carlo-simulation/paytable"
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
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(reels[i])-1)))
		if err != nil {
			log.Fatal(err)
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

func transpose(a [][]int) [][]int {
	newArr := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			newArr[j] = append(newArr[j], a[i][j])
		}
	}

	return newArr
}

func (m *Matrix) CheckForFreeGames(scatter int) int {
	transposedMatrix := transpose(m.Matrix)
	numOfScatters := 0
	for i := range transposedMatrix {
		for j := range transposedMatrix[i] {
			if scatter == transposedMatrix[i][j] {
				numOfScatters++
				break
			}
		}
	}

	return paytable.Desc.FreeGames[numOfScatters]
}

func (m *Matrix) getScatterPayoff(scatter int) int {
	transposedMatrix := transpose(m.Matrix)
	numOfScatters := 0
	for i := range transposedMatrix {
		for j := range transposedMatrix[i] {
			if scatter == transposedMatrix[i][j] {
				numOfScatters++
				break
			}
		}
	}

	return paytable.Desc.Paytable[scatter][numOfScatters]
}
