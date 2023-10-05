package reels

import (
	"fmt"
	"log"
	"monte-carlo-simulation/gameDescription"

	"gonum.org/v1/gonum/mat"
)

func GetCombinationWeights() {
	var coeffArr []float64
	for i := range gameDescription.Desc.Paytable {
		for j := range gameDescription.Desc.Paytable[i] {
			coeffArr = append(coeffArr, float64(gameDescription.Desc.Paytable[i][j]))
		}
	}
	coefficients := mat.NewDense(1, len(coeffArr), coeffArr)
	constants := mat.NewVecDense(1, []float64{0.95})
	variables := mat.NewVecDense(len(coeffArr), nil)

	err := variables.SolveVec(coefficients, constants)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Solution: %v", variables)
}
