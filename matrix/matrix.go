package matrix

import (
	"crypto/rand"
	"log"
	"math/big"
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

// func (m *Matrix) checkForFreeGames(scatter int) int {

// }
