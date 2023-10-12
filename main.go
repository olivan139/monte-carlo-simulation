package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"monte-carlo-simulation/model"
	"monte-carlo-simulation/simulation"

	"github.com/joho/godotenv"
)

func main() {
	if _, err := os.Stat("data/logs/"); os.IsNotExist(err) {
		if err = os.Mkdir("data/logs", os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
	file, err := os.OpenFile("data/logs/logs "+string(time.Now().Format("2006-01-02T15:04:05"))+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
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

	model.ParseDescriptionJSON(filePath)

	start := time.Now()
	simulation.StartSimulation(10000000)
	elapsed := time.Since(start)
	fmt.Printf("Time passed: %v\n", elapsed)
	log.Fatal("Application finished successfully")
}
