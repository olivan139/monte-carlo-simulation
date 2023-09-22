package main

import (
	"fmt"
	"log"
	"os"

	"monte-carlo-simulation/matrix"
	"monte-carlo-simulation/paytable"

	"github.com/joho/godotenv"
)

func main() {
	file, err := os.OpenFile("logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	log.Println("application has started")

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	paytable.ReadFromFile(os.Getenv("PAYTABLE_FILE_PATH"))

	var matrix matrix.Matrix
	matrix.Init(3, 3)
	reels := [][]int{
		[]int{1, 2, 3, 4, 5, 6, 6, 6, 6, 7, 8},
		[]int{1, 2, 3, 4, 5, 6, 6, 6, 6, 7, 8},
		[]int{1, 2, 3, 4, 5, 6, 6, 6, 6, 7, 8},
	}

	err = matrix.GenerateFromReels(reels)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(paytable.Paytable)
	fmt.Println(matrix.Matrix)
}
