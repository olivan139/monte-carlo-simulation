package main

import (
	"log"
	"os"
	"time"

	"monte-carlo-simulation/gameDescription"
	"monte-carlo-simulation/simulation"

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

	// reels1 := [][]int{
	// 	{6, 8, 8, 3, 4, 5, 2, 11, 5, 4, 2, 3, 2, 4, 6, 5, 2, 6, 8, 6, 5, 7, 4, 9, 2, 11, 3, 3, 6, 2, 6, 5, 3, 4, 7, 9, 4, 10, 5, 8, 2, 3, 7, 3, 10, 11, 2, 2, 4, 5, 7, 2, 10, 2, 3, 4, 4, 7, 3, 9, 4, 5, 6, 4, 3, 6, 2, 7, 2, 5, 2, 8, 3, 5, 2, 4, 3, 9, 6, 2, 8, 4, 2, 6, 3, 9, 3},
	// 	{8, 2, 4, 8, 4, 5, 9, 3, 7, 3, 4, 11, 3, 3, 4, 2, 6, 2, 3, 7, 8, 3, 2, 5, 2, 4, 5, 2, 3, 6, 6, 5, 4, 4, 8, 2, 9, 3, 3, 7, 2, 7, 5, 5, 4, 2, 5, 5, 6, 9, 7, 5, 9, 11, 3, 4, 6, 2, 3, 2, 5, 10, 4, 2, 10, 9, 4, 8, 2, 5, 2, 10, 2, 3, 6, 6, 4, 11, 2, 8, 3, 2, 6},
	// 	{10, 1, 3, 7, 5, 11, 9, 1, 2, 6, 4, 8, 10, 7, 4, 4, 8, 6, 3, 9, 5, 8, 3, 2, 4, 2, 3, 3, 5, 2, 4, 5, 5, 1, 3, 11, 6, 2, 1, 2, 4, 6, 2, 9, 2, 8, 2, 5, 4, 3, 6, 4, 4, 2, 5, 7, 4, 9, 3, 4, 4, 7, 6, 6, 10, 3, 8, 7, 3, 3, 5, 3, 5, 3, 3, 7, 3, 6, 5, 10, 2, 4, 2, 2, 4, 11, 2, 2, 5, 6, 4, 2, 9},
	// 	{2, 6, 2, 10, 2, 6, 6, 11, 2, 4, 8, 9, 2, 3, 3, 7, 2, 5, 8, 5, 5, 10, 7, 6, 3, 3, 11, 10, 4, 3, 4, 2, 2, 3, 5, 2, 2, 5, 5, 6, 4, 2, 5, 9, 6, 8, 3, 9, 3, 3, 5, 9, 3, 8, 11, 4, 6, 3, 4, 3, 6, 9, 2, 7, 5, 8, 6, 2, 3, 2, 4, 4, 4, 7, 4, 7, 2, 4, 6, 8, 7, 2, 4, 3, 5, 2, 2},
	// 	{2, 5, 2, 2, 8, 6, 8, 8, 3, 4, 8, 6, 7, 6, 4, 6, 6, 5, 3, 3, 4, 11, 2, 2, 3, 3, 4, 3, 4, 3, 7, 5, 7, 5, 2, 4, 2, 7, 3, 4, 4, 4, 2, 8, 9, 9, 3, 4, 2, 4, 5, 3, 6, 7, 2, 2, 5, 4, 3, 11, 3, 4, 9, 10, 2, 5, 8, 11, 2, 6, 10, 6, 10, 5, 2, 3, 2, 5, 5, 2, 6, 2, 4, 5, 2, 2, 10, 3, 9},
	// }
	// var mat matrix.Matrix
	// mat.Init(3, 5)
	// for i := 0; i < 20; i++ {
	// 	mat.GenerateFromReels(reels1)
	// 	for _, val := range mat.Matrix {
	// 		fmt.Println(val)
	// 	}
	// 	spinWin := 0
	// 	for _, line := range gameDescription.Desc.Lines {
	// 		winLine := []int{mat.Matrix[line[0]][0], mat.Matrix[line[1]][1], mat.Matrix[line[2]][2], mat.Matrix[line[3]][3], mat.Matrix[line[4]][4]}
	// 		spinWin += matrix.GetLinePayoff(winLine)
	// 	}

	// 	fmt.Println(spinWin)
	// }

	simulation.StartSimulation()

	log.Fatal("Application finished successfully")
}
