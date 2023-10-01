package helper

import (
	"golang.org/x/exp/constraints"
)

func Transpose(slice [][]int) [][]int {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]int, xl)
	for i := range result {
		result[i] = make([]int, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func Max[T constraints.Ordered](el1 T, el2 T) T {
	if el1 > el2 {
		return el1
	}

	return el2
}
