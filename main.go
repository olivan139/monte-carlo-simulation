package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"monte-carlo-simulation/gameDescription"
	"monte-carlo-simulation/matrix"

	"github.com/joho/godotenv"
)

func main() {
	if _, err := os.Stat("logs/"); os.IsNotExist(err) {
		if err = os.Mkdir("logs", os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
	file, err := os.OpenFile("logs/logs "+string(time.Now().Format("2006-01-02T15:04:05"))+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Panic(err)
	}

	log.SetOutput(file)

	log.Println("application has started")

	if err := godotenv.Load(); err != nil {
		log.Panic(err)
	}

	filePath := os.Getenv("PAYTABLE_FILE_PATH")
	if filePath == "" {
		log.Panic("No file path was given, check .env file\n")
	}

	gameDescription.ParseDescriptionJSON(filePath)

	var mat matrix.Matrix
	mat.Init(3, 3)
	reels := [][]int{
		{1, 2, 3, 4, 5, 6, 6, 6, 6, 7, 8},
		{1, 2, 3, 4, 5, 6, 6, 6, 6, 7, 8},
		{1, 2, 3, 4, 5, 6, 6, 6, 6, 7, 8},
	}

	err = mat.GenerateFromReels(reels)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(gameDescription.Desc.Paytable)
	fmt.Println(mat.Matrix)

	fmt.Println(matrix.GetLinePayoff([]int{0, 0, 0, 0, 0}))
	fmt.Println(matrix.GetLinePayoff([]int{10, 10, 0, 0, 0}))
	fmt.Println(matrix.GetLinePayoff([]int{10, 10, 10, 10, 0}))
	fmt.Println(matrix.GetLinePayoff([]int{0, 1, 2, 3, 4}))
	fmt.Println(matrix.GetLinePayoff([]int{3, 3, 10, 4, 3}))
	fmt.Println(matrix.GetLinePayoff([]int{5, 5, 5, 5, 10}))
	fmt.Println(matrix.GetLinePayoff([]int{10, 5, 5, 4, 0}))
	fmt.Println(matrix.GetLinePayoff([]int{5, 5, 5, 5, 5}))
	fmt.Println(matrix.GetLinePayoff([]int{10, 10, 10, 10, 10}))

	log.Fatal("Application finished successfully")
}
