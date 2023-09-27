package helper

import (
	"golang.org/x/exp/constraints"
)

func Transpose(a [][]int) [][]int {
	newArr := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			newArr[j] = append(newArr[j], a[i][j])
		}
	}

	return newArr
}

func Max[T constraints.Ordered](el1 T, el2 T) T {
	if el1 > el2 {
		return el1
	}

	return el2
}
